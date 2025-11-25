package firebase

import (
	"context"

	fbauth "firebase.google.com/go/v4/auth"
	"github.com/radovskyb/scaff-firebase-auth/auth"
)

type client struct {
	cl *fbauth.Client
}

func NewClient(fbclient *fbauth.Client) auth.Client {
	return &client{cl: fbclient}
}

func (c *client) VerifyIDToken(ctx context.Context, token string) (*auth.Token, error) {
	decodedToken, err := c.cl.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	// Create our auth.Token type from the decodedToken info.
	t := &auth.Token{
		UID: decodedToken.UID,
	}

	// Let's check Firebase claims to grab the email and display name.
	if name, found := decodedToken.Claims["name"].(string); found {
		t.DisplayName = name
	}
	if email, found := decodedToken.Claims["email"].(string); found {
		t.EmailAddress = email
	}

	return t, nil
}
