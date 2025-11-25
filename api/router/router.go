package router

import (
	"encoding/json"
	"net/http"

	"github.com/radovskyb/scaff-firebase-auth/cfg"
	"github.com/radovskyb/scaff-firebase-auth/handler"
)

func Setup(dp *cfg.APIDependencies) (http.Handler, error) {
	h := handler.New(dp.AuthClient)

	f := handler.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return json.NewEncoder(w).Encode(map[string]int{"status": http.StatusOK})
	})

	mux := http.NewServeMux()
	mux.Handle("/api/protected", CorsMiddleware(h.Serve(f)))

	return mux, nil
}

// CorsMiddleware sets the required cors headers for our API.
//
// You'll want to create a proper cors middleware, but just so this works during dev for the example,
// I'm setting ACAO to localhost:5173 directly
func CorsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,OPTIONS,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
