package cfg

import (
	"context"

	fbauth "firebase.google.com/go/v4"
	"github.com/radovskyb/scaff-firebase-auth/auth"
	"github.com/radovskyb/scaff-firebase-auth/auth/firebase"
	"google.golang.org/api/option"
)

type APIDependencies struct {
	AuthClient auth.Client
}

func LoadDeps(ctx context.Context, envCfg *Env) (*APIDependencies, error) {
	sa := option.WithCredentialsFile(envCfg.ServiceKeyFile)
	app, err := fbauth.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	fbauthClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	// Here we create our auth client from the firebase client.
	authClient := firebase.NewClient(fbauthClient)

	ad := &APIDependencies{
		AuthClient: authClient,
	}

	return ad, nil
}
