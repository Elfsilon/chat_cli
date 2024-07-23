package utils

import (
	"errors"
	"server/pkg/utils/role"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrMethodNotFound   = errors.New("method not found")
	ErrMetadataMissing  = errors.New("metadata is missing")
	ErrTokenMissing     = errors.New("token is missing")
	ErrInvalidStructure = errors.New("auth header structure is invalid")
	ErrNoPermission     = errors.New("no permission")
)

type RequestAuthorizer struct {
	tm      *TokenManager
	options map[string]*MethodOptions
}

func NewRequestAuthorizer(tm *TokenManager) *RequestAuthorizer {
	opt := map[string]*MethodOptions{
		"/auth.Auth/Login":           NewMethodOptions().SetProtected(false),
		"/auth.Auth/GetRefreshToken": NewMethodOptions().SetProtected(false),
		"/auth.Auth/GetAccessToken":  NewMethodOptions().SetProtected(false),
		"/auth.Auth/CheckResource":   NewMethodOptions().SetProtected(false),

		"/user.UserService/Create":  NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),
		"/user.UserService/Get":     NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),
		"/user.UserService/GetList": NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),
		"/user.UserService/Update":  NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),
		"/user.UserService/Delete":  NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),

		"/chat.Chat/Create":      NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),
		"/chat.Chat/Conect":      NewMethodOptions().SetProtected(true).SetAllowedRoles(role.User, role.Admin),
		"/chat.Chat/SendMessage": NewMethodOptions().SetProtected(true).SetAllowedRoles(role.User, role.Admin),
		"/chat.Chat/Delete":      NewMethodOptions().SetProtected(true).SetAllowedRoles(role.Admin),
	}

	return &RequestAuthorizer{
		tm:      tm,
		options: opt,
	}
}

func (a *RequestAuthorizer) parseTokenHeader(tokens []string, prefix string) (Claims, error) {
	if len(tokens) == 0 || tokens[0] == "" {
		return errorClaims, status.Errorf(codes.Unauthenticated, ErrTokenMissing.Error())
	}

	parts := strings.Split(tokens[0], " ")
	if len(parts) != 2 && parts[0] != prefix {
		return errorClaims, status.Errorf(codes.Unauthenticated, ErrInvalidStructure.Error())
	}

	token := parts[1]
	claims, err := a.tm.GetAccessTokenClaims(token)
	if err != nil {
		return errorClaims, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return claims, nil
}

// "authorization: Bearer <token>" - must be provided with metadata if method is protected
func (a *RequestAuthorizer) AuthorizeUser(md metadata.MD, method string) (Claims, error) {
	opt, ok := a.options[method]
	if !ok {
		return errorClaims, status.Errorf(codes.NotFound, ErrMethodNotFound.Error())
	}

	if !opt.protected {
		return errorClaims, nil
	}

	claims, err := a.parseTokenHeader(md["authorization"], "Bearer")
	if err != nil {
		return errorClaims, err
	}

	if !opt.IsAllowedForRole(claims.Role) {
		return errorClaims, status.Errorf(codes.PermissionDenied, ErrNoPermission.Error())
	}

	return claims, nil
}

func (a *RequestAuthorizer) AuthorizeService(md metadata.MD) error {
	_, err := a.parseTokenHeader(md["s-authorization"], "Service")
	return err
}
