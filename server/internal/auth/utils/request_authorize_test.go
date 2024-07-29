package utils

import (
	"fmt"
	"server/pkg/utils/role"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestUnprotectedWithTokenIsOk(t *testing.T) {
	allAllowedMethod := "/auth.Auth/Login"
	claims, err := helper(t, allAllowedMethod, false, role.User)
	assert.NoError(t, err)
	assert.Equal(t, Claims{}, claims)
}

func TestPotectedWithTokenIsAborted(t *testing.T) {
	adminPermissionsMethod := "/user.UserService/Create"
	claims, err := helper(t, adminPermissionsMethod, false, role.User)
	assert.Error(t, err)
	assert.Equal(t, Claims{}, claims)
}

func TestPotectedWithUserTokenIsOk(t *testing.T) {
	userPermissionsMethod := "/chat.Chat/List"
	claims, err := helper(t, userPermissionsMethod, true, role.User)
	assert.NoError(t, err)
	assert.NotEqual(t, Claims{}, claims)
}

func TestPotectedWithUserTokenIsAborted(t *testing.T) {
	adminPermissionsMethod := "/user.UserService/Create"
	claims, err := helper(t, adminPermissionsMethod, true, role.User)
	assert.Error(t, err)
	assert.Equal(t, Claims{}, claims)
}

func helper(t *testing.T, method string, tokenEnabled bool, role role.Role) (Claims, error) {
	t.Helper()

	ttl := 5 * time.Second
	tm := NewTokenManager(ttl, ttl, []byte("test"))
	a := NewRequestAuthorizer(tm)

	md := metadata.Pairs()
	if tokenEnabled {
		token, err := tm.GenerateAccessToken(0, "John", role)
		if err != nil {
			return Claims{}, err
		}
		md = metadata.Pairs("authorization", fmt.Sprintf("Bearer %v", token))
	}
	return a.AuthorizeUser(md, method)
}
