package services

import (
	"context"
	"fmt"
	"server/internal/auth/gen/user"
	"server/internal/auth/mocks"
	"server/internal/auth/repos"
	"server/internal/auth/utils"
	"server/pkg/utils/role"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

const (
	freeResource             = "/auth.Auth/Login"
	userPermissionsResource  = "/chat.Chat/List"
	adminPermissionsResource = "/user.UserService/GetList"
)

func initialize(t *testing.T, tokenized bool, r role.Role) (context.Context, *AuthService, error) {
	t.Helper()

	ttl := 3 * time.Second
	tm := utils.NewTokenManager(ttl, ttl, []byte("test"))
	tu := utils.NewTokenUpdater(tm, ttl)
	ra := utils.NewRequestAuthorizer(tm)

	rep := repos.NewSessionMockRepo()
	if _, err := rep.Create(context.Background(), 1, "ref", time.Now().Add(time.Hour)); err != nil {
		return nil, nil, err
	}

	client := mocks.NewUserServiceMockClient()
	req := &user.CreateUserRequest{
		Name:     "Julian",
		Email:    "",
		Password: "$2a$14$f925wgPQFT7SO1dIo9.bd.DB8UZzCzTVAuypqp0xZkvBP9lDJebnC",
		Role:     user.Role_User,
	}
	if _, err := client.Create(context.Background(), req); err != nil {
		return nil, nil, err
	}

	s := NewAuthService("", client, rep, tm, tu, ra)

	ctx := context.Background()
	if tokenized {
		token, err := tm.GenerateAccessToken(0, "Julian", r)
		if err != nil {
			return nil, nil, err
		}
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", fmt.Sprintf("Bearer %v", token)))
	} else {
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("t", "t"))
	}

	return ctx, s, nil
}

func TestCheckUnprotectedResourceGrantedWithoutToken(t *testing.T) {
	ctx, c, err := initialize(t, false, role.User)
	assert.NoError(t, err)

	_, err = c.CheckAccess(ctx, freeResource)
	assert.NoError(t, err)
}

func TestCheckProtectedResourceDeniedWithoutToken(t *testing.T) {
	ctx, c, err := initialize(t, false, role.User)
	assert.NoError(t, err)

	_, err = c.CheckAccess(ctx, userPermissionsResource)
	assert.Error(t, err)
}

func TestCheckUserResourceWithUserPermissionsGranted(t *testing.T) {
	ctx, c, err := initialize(t, true, role.User)
	assert.NoError(t, err)

	_, err = c.CheckAccess(ctx, userPermissionsResource)
	assert.NoError(t, err)
}

func TestCheckAdminResourceWithUserPermissionsDenied(t *testing.T) {
	ctx, c, err := initialize(t, true, role.User)
	assert.NoError(t, err)

	_, err = c.CheckAccess(ctx, adminPermissionsResource)
	assert.Error(t, err)
}

func TestCheckAdminResourceWithAdminPermissionsGranted(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	_, err = c.CheckAccess(ctx, adminPermissionsResource)
	assert.NoError(t, err)
}

func TestGetAccessToken(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	token, err := c.GetAccessToken(ctx, "ref")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGetRefreshToken(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	token, err := c.GetRefreshToken(ctx, "ref")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEqual(t, "ref", token)
}

func TestLogin(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	_, token, err := c.Login(ctx, "Julian", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
