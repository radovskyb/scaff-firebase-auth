package main

import (
	"context"
	"log"
	"net/http"

	"github.com/radovskyb/scaff-firebase-auth/cfg"
	"github.com/radovskyb/scaff-firebase-auth/router"
)

func main() {
	ctx := context.Background()

	// Load our env config.
	envCfg, err := cfg.LoadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	dp, err := cfg.LoadDeps(ctx, envCfg)
	if err != nil {
		log.Fatalln(err)
	}

	// Fetch our api routes, passing our dependencies to our router.
	r, err := router.Setup(dp)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(http.ListenAndServe(envCfg.PortStr, r))
}
