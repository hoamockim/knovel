package auth

import "context"

type Authentication interface {
	SignIn(ctx context.Context, req *SignInRequest) (*ClaimInfo, error)
}
