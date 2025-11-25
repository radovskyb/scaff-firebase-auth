package auth

import "context"

// Client is our custom auth client interface. Our implementation will be using Firebase.
type Client interface {
	VerifyIDToken(ctx context.Context, token string) (*Token, error)
}

type Token struct {
	UID          string
	EmailAddress string
	DisplayName  string
}
