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
	mux.Handle("/api", h.Serve(f))

	return mux, nil
}
