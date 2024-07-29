package controllers

import (
	"context"
	"fmt"
	"server/internal/auth/gen/user"
	"server/internal/auth/mocks"
	"server/internal/auth/repos"
	"server/internal/auth/services"
	"server/internal/auth/utils"
	"server/pkg/gen/auth"
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

func initialize(t *testing.T, tokenized bool, r role.Role) (context.Context, *AuthController, error) {
	t.Helper()

	ttl := 3 * time.Second
	tm := utils.NewTokenManager(ttl, ttl, []byte("test"))
	tu := utils.NewTokenUpdater(tm, ttl)
	ra := utils.NewRequestAuthorizer(tm)

	rep := repos.NewSessionMockRepo()
	if _, err := rep.Create(context.Background(), 0, "ref", time.Now().Add(time.Hour)); err != nil {
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

	s := services.NewAuthService("", client, rep, tm, tu, ra)
	c := NewAuthController(s)

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

	return ctx, c, nil
}

func TestCheckUnprotectedResourceGrantedWithoutToken(t *testing.T) {
	ctx, c, err := initialize(t, false, role.User)
	assert.NoError(t, err)

	req := &auth.CheckResourceRequest{FullMethod: freeResource}
	res, _ := c.CheckResource(ctx, req)
	assert.Empty(t, "", res.Reason)
	assert.Equal(t, true, res.HasAccess)
}

func TestCheckProtectedResourceDeniedWithoutToken(t *testing.T) {
	ctx, c, err := initialize(t, false, role.User)
	assert.NoError(t, err)

	req := &auth.CheckResourceRequest{FullMethod: userPermissionsResource}
	res, _ := c.CheckResource(ctx, req)
	assert.NotEmpty(t, res.Reason)
	assert.Equal(t, false, res.HasAccess)
}

func TestCheckUserResourceWithUserPermissionsGranted(t *testing.T) {
	ctx, c, err := initialize(t, true, role.User)
	assert.NoError(t, err)

	req := &auth.CheckResourceRequest{FullMethod: userPermissionsResource}
	res, _ := c.CheckResource(ctx, req)
	assert.Empty(t, res.Reason)
	assert.Equal(t, true, res.HasAccess)
}

func TestCheckAdminResourceWithUserPermissionsDenied(t *testing.T) {
	ctx, c, err := initialize(t, true, role.User)
	assert.NoError(t, err)

	req := &auth.CheckResourceRequest{FullMethod: adminPermissionsResource}
	res, _ := c.CheckResource(ctx, req)
	assert.NotEmpty(t, res.Reason)
	assert.Equal(t, false, res.HasAccess)
}

func TestCheckAdminResourceWithAdminPermissionsGranted(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	req := &auth.CheckResourceRequest{FullMethod: adminPermissionsResource}
	res, _ := c.CheckResource(ctx, req)
	assert.Empty(t, res.Reason)
	assert.Equal(t, true, res.HasAccess)
}

func TestGetAccessToken(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	req := &auth.GetAccessTokenRequest{RefreshToken: "ref"}
	res, err := c.GetAccessToken(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.AccessToken)
}

func TestGetRefreshToken(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	req := &auth.GetRefreshTokenRequest{RefreshToken: "ref"}
	res, err := c.GetRefreshToken(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.RefreshToken)
	assert.NotEqual(t, "ref", res.RefreshToken)
}

func TestLogin(t *testing.T) {
	ctx, c, err := initialize(t, true, role.Admin)
	assert.NoError(t, err)

	req := &auth.LoginRequest{Name: "Julian", Password: "password"}
	res, err := c.Login(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.RefreshToken)
}
