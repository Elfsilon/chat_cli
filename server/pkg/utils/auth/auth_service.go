package authutils

import (
	"context"
	"server/pkg/gen/auth"
	"server/pkg/utils/hasher"
)

func GetServiceToken(ctx context.Context, authService auth.AuthClient, secret string) (string, error) {
	secretHash, err := hasher.HashPassword(secret)
	if err != nil {
		return "", err
	}

	req := &auth.GetServiceTokenRequest{Hash: secretHash}
	resp, err := authService.GetServiceToken(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
