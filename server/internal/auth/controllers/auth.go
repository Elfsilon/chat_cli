package controllers

import (
	"context"
	"server/internal/auth/services"
	"server/pkg/gen/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	authService *services.AuthService
	auth.UnimplementedAuthServer
}

func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{
		authService: as,
	}
}

func (c *AuthController) HealthCheck(ctx context.Context, req *auth.HealthRequest) (*auth.HealthResponse, error) {
	return &auth.HealthResponse{}, nil
}

func (c *AuthController) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	_, refreshToken, err := c.authService.Login(ctx, req.GetName(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	reponse := &auth.LoginResponse{
		RefreshToken: refreshToken,
	}

	return reponse, nil
}

func (c *AuthController) GetRefreshToken(ctx context.Context, req *auth.GetRefreshTokenRequest) (*auth.GetRefreshTokenResponse, error) {
	refreshToken, err := c.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	reponse := &auth.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}

	return reponse, nil
}

func (c *AuthController) GetAccessToken(ctx context.Context, req *auth.GetAccessTokenRequest) (*auth.GetAccessTokenResponse, error) {
	accessToken, err := c.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	reponse := &auth.GetAccessTokenResponse{
		AccessToken: accessToken,
	}

	return reponse, nil
}

func (c *AuthController) GetServiceToken(ctx context.Context, req *auth.GetServiceTokenRequest) (*auth.GetServiceTokenResponse, error) {
	accessToken, err := c.authService.GetServiceToken(ctx, req.GetHash())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	reponse := &auth.GetServiceTokenResponse{
		AccessToken: accessToken,
	}

	return reponse, nil
}

func (c *AuthController) CheckResource(ctx context.Context, req *auth.CheckResourceRequest) (*auth.CheckResourceResponse, error) {
	userID, err := c.authService.CheckAccess(ctx, req.GetFullMethod())

	response := auth.CheckResourceResponse{}
	if err != nil {
		response.UserID = -1 // TODO:
		response.Reason = err.Error()
	} else {
		response.HasAccess = true
		response.UserID = int64(userID)
	}

	return &response, nil
}
