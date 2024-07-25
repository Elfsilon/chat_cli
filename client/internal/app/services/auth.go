package services

import (
	"chat_cli/internal/app/gen/auth"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrTokenExpiredOrEmpty = errors.New("token is expired or empty")
	ErrSessionNotFound     = errors.New("session is not found")
)

type Claims struct {
	UserID int    `json:"uid"`
	Name   string `json:"name"`
	Role   int    `json:"role"`
	jwt.StandardClaims
}

const stateFilePath = "/Users/user/Desktop/Test/client/state.json"

type AuthService struct {
	client       auth.AuthClient
	ttl          time.Duration
	secret       []byte
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	claims       Claims
}

func NewAuthService(client auth.AuthClient, ttl time.Duration, secret []byte) *AuthService {
	return &AuthService{client: client, ttl: ttl, secret: secret}
}

func (s *AuthService) Context(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %v", s.AccessToken))
}

func (s *AuthService) GetClaims() Claims {
	return s.claims
}

func (s *AuthService) Load() error {
	bytes, err := os.ReadFile(stateFilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, s)
}

func (s *AuthService) Save() error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(stateFilePath, bytes, os.ModeAppend)
}

func (s *AuthService) Login(ctx context.Context, name, password string) error {
	req := &auth.LoginRequest{
		Name:     name,
		Password: password,
	}

	res, err := s.client.Login(ctx, req)
	if err != nil {
		return err
	}
	s.RefreshToken = res.GetRefreshToken()

	return nil
}

func (s *AuthService) UpdateRefreshToken(ctx context.Context) error {
	if s.RefreshToken == "" {
		return ErrTokenExpiredOrEmpty
	}

	req := &auth.GetRefreshTokenRequest{
		RefreshToken: s.RefreshToken,
	}

	res, err := s.client.GetRefreshToken(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.Unauthenticated {
				return ErrTokenExpiredOrEmpty
			}
			if st.Code() == codes.Internal {
				if st.Message() == "ent: session not found" {
					return ErrSessionNotFound
				}
			}
		}
		return err
	}
	s.RefreshToken = res.GetRefreshToken()

	return nil
}

func (s *AuthService) updateAccessTokenClaims() error {
	token, err := jwt.ParseWithClaims(s.AccessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return errors.New("invalid claims")
	}

	s.claims = *claims
	return nil
}

func (s *AuthService) UpdateAccessToken(ctx context.Context) error {
	req := &auth.GetAccessTokenRequest{
		RefreshToken: s.RefreshToken,
	}

	res, err := s.client.GetAccessToken(ctx, req)
	if err != nil {
		return err
	}

	s.AccessToken = res.GetAccessToken()
	return s.updateAccessTokenClaims()
}

func (s *AuthService) RunAccessTokenUpdater(ctx context.Context) {
	t := time.NewTicker(s.ttl)
	for {
		select {
		case <-t.C:
			if err := s.UpdateAccessToken(ctx); err != nil {
				fmt.Printf("failed updating access token: %v\n", err)
			}
		case <-ctx.Done():
			return
		}
	}
}
