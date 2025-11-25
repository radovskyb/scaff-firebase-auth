package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/radovskyb/scaff-firebase-auth/auth"
)

type Handler struct {
	ac auth.Client
}

// New accepts an auth.Client interface. This is what we're going to pass to our auth middleware function in Server.
func New(ac auth.Client) *Handler {
	return &Handler{ac: ac}
}

// HandlerFunc is the custom impl that we'll use for Serve.
//
// This is a personal preference, but I've always loved creating a custom HandlerFunc that returns an error,
// since it makes errors handling an easy process when all errors propagate up the handler/middleware chain into Serve.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Serve is our main handler
func (h *Handler) Serve(next HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx, secure, err := AuthCheck(r, h.ac)
		if !secure {
			h.writeError(w, err)
			return
		}

		err = next(w, r.WithContext(rctx))
		if err == nil {
			return
		}

		h.writeError(w, err)
	})
}

// Err is our custom type that we'll use to encode as JSON when we write errors.
type Err struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

// writeError is where I like to handle any errors that were returned throughout Server and it's handler funcs.
//
// In this implementation, I'm accepting the ResponseWriter and an error, but it's a pretty good idea to either accept a env flag
// to distinguish between dev and prod since logging and output shouldn't be the same for security reasons.
//
// You can also keep vars like Env directly within your Handler struct and pass them to NewHandler.
func (h *Handler) writeError(w http.ResponseWriter, err error) {
	// Quick example:
	//
	// Default to internal server error, so unknown errors don't leak to users in the frontend.
	status := http.StatusInternalServerError
	msg := "An unknown error has occured"

	// Check what type of error is being returned.
	if errors.Is(err, ErrAuthCheck) || errors.Is(err, ErrInvalidAuthHeaderFormat) {
		status, msg = http.StatusUnauthorized, err.Error()
	}

	w.WriteHeader(status)

	// Create our error to send to the client.
	//
	// Note: If we don't end up checking for specific errors that don't get handled above,
	// this is where the default 500 and unknown error message will be used.
	jsonErr := &Err{Msg: msg, Status: status}

	json.NewEncoder(w).Encode(jsonErr)
}
