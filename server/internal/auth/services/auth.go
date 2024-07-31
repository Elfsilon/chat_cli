package services

import (
	"context"
	"errors"
	"fmt"
	"server/internal/auth/gen/user"
	"server/internal/auth/repos"
	"server/internal/auth/utils"
	"server/pkg/utils/hasher"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidCreds     = errors.New("invalid credentials")
	ErrSessionIsExpired = errors.New("session is expired, refresh token need to be updated")
	ErrAuthFailed       = errors.New("authorization failed")
	ErrMetadataMissing  = errors.New("metadata is missing")
)

type Auth interface {
	CheckAccess(ctx context.Context, method string) (int, error)
	CheckToken(ctx context.Context, token string) error
	Login(ctx context.Context, name string, password string) (*int, string, error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetServiceToken(ctx context.Context, secretHash string) (string, error)
}

type AuthService struct {
	serviceSecret string
	uclient       user.UserServiceClient
	sessionRepo   repos.Session
	tokenManager  *utils.TokenManager
	tokenUpdater  *utils.TokenUpdater
	authorizer    *utils.RequestAuthorizer
}

func NewAuthService(
	serviceSecret string,
	us user.UserServiceClient,
	sr repos.Session,
	tm *utils.TokenManager,
	tu *utils.TokenUpdater,
	ra *utils.RequestAuthorizer,
) *AuthService {
	return &AuthService{
		serviceSecret: serviceSecret,
		uclient:       us,
		sessionRepo:   sr,
		tokenManager:  tm,
		tokenUpdater:  tu,
		authorizer:    ra,
	}
}

// If requester hasn't access to the resource then returns an error
func (s *AuthService) CheckAccess(ctx context.Context, method string) (int, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return -1, status.Errorf(codes.InvalidArgument, ErrMetadataMissing.Error())
	}

	claims, err := s.authorizer.AuthorizeUser(md, method)
	if err != nil {
		if errs := s.authorizer.AuthorizeService(md); errs != nil {
			return -1, fmt.Errorf("unable to authorize as a user: %v\nunable to authorize as a service: %v", err, errs)
		}
		// This means that method was requested by authorized service with full access
		return -1, nil
	}
	return claims.UserID, nil
}

func (s *AuthService) CheckToken(ctx context.Context, token string) error {
	_, err := s.tokenManager.GetAccessTokenClaims(token)
	return err
}

func (s *AuthService) Login(ctx context.Context, name string, password string) (*int, string, error) {
	ctx = s.tokenUpdater.Context(ctx)
	req := &user.GetRequest{Username: name}
	u, err := s.uclient.Get(ctx, req)
	if err != nil {
		return nil, "", err
	}

	passed := hasher.CheckPasswordHash(password, u.GetPasswordHash())
	if !passed {
		return nil, "", ErrInvalidCreds
	}

	refreshToken, expiresAt := s.tokenManager.GenerateRefreshToken()
	sessionID, err := s.sessionRepo.Create(ctx, int(u.Id), refreshToken, expiresAt)
	if err != nil {
		return nil, "", err
	}

	return &sessionID, refreshToken, nil
}

func (s *AuthService) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	session, err := s.sessionRepo.GetByRefreshToken(ctx, oldRefreshToken)
	if err != nil {
		return "", err
	}

	refreshToken, expiresAt := s.tokenManager.GenerateRefreshToken()
	if err := s.sessionRepo.UpdateRefreshToken(ctx, session.ID, refreshToken, expiresAt); err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *AuthService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	session, err := s.sessionRepo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	sessionIsExpired := time.Now().UTC().After(session.ExpiredAt)
	if sessionIsExpired {
		return "", ErrSessionIsExpired
	}

	ctx = s.tokenUpdater.Context(ctx)
	req := &user.GetRequest{Id: int64(session.UserID)}
	user, err := s.uclient.Get(ctx, req)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user is nil")
	}
	accessToken, err := s.tokenManager.GenerateAccessToken(session.UserID, user.Name, int(user.Role))
	if err != nil {
		return "", nil
	}

	return accessToken, nil
}

func (s *AuthService) GetServiceToken(ctx context.Context, secretHash string) (string, error) {
	if !hasher.CheckPasswordHash(s.serviceSecret, secretHash) {
		return "", ErrAuthFailed
	}

	accessToken, err := s.tokenManager.GenerateServiceToken()
	if err != nil {
		return "", nil
	}

	return accessToken, nil
}
