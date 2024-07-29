package utils

import (
	"server/pkg/utils/role"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTM(t *testing.T) {
	ttl := 5 * time.Second
	tm := NewTokenManager(ttl, ttl, []byte("test"))

	token, err := tm.GenerateAccessToken(0, "John", role.User)
	assert.NoError(t, err)

	claims, err := tm.GetAccessTokenClaims(token)
	assert.NoError(t, err)

	assert.Equal(t, 0, claims.UserID)
	assert.Equal(t, "John", claims.Name)
	assert.Equal(t, role.User, claims.Role)
}
