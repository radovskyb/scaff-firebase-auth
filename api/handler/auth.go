package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/radovskyb/scaff-firebase-auth/auth"
)

var (
	ErrInvalidAuthHeaderFormat = errors.New("invalid Authorization header format (Bearer token expected)")
	ErrAuthCheck               = errors.New("jwt invalid")
)

func tokenAuthCheck(ctx context.Context, ac auth.Client, idToken string) (*auth.Token, error) {
	decodedToken, err := ac.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, ErrAuthCheck
	}

	return decodedToken, nil
}

// It's recommended to create a custom context key type.
type ctxKey string

const decodedTokenCtxKey ctxKey = "decodedToken"

// AuthCheck reads a request's Authorization header and attempts to fetch
func AuthCheck(r *http.Request, ac auth.Client) (context.Context, bool, error) {
	var idToken string

	authHeader := r.Header.Get("Authorization")

	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		idToken = parts[1]
	} else {
		return r.Context(), false, ErrInvalidAuthHeaderFormat
	}

	decodedToken, err := tokenAuthCheck(r.Context(), ac, idToken)
	if err != nil {
		return r.Context(), false, err
	}

	// This is a good spot to fetch or create a user in your store based on the decodedToken.
	// e.g user, err := h.fetchOrCreateUser(decodedToken)
	// ...
	//
	// ctx := context.WithValue(r.Context(), FetchUserContextKey, user)
	// return ctx, true, nil

	// For now, simply store the decodedToken for use in wrapper handlers functions.
	ctx := context.WithValue(r.Context(), decodedTokenCtxKey, decodedToken)

	return ctx, true, nil
}

// GetDecodedTokenFromContext will be callable from within the handler package, for when we
// want to fetch and use info from the token.
func GetDecodedTokenFromContext(ctx context.Context) (*auth.Token, bool) {
	token, ok := ctx.Value(decodedTokenCtxKey).(*auth.Token)
	if !ok || token == nil {
		return nil, false
	}
	return token, true
}
