package services

import (
	"context"
	"server/internal/user/repos"
	"server/pkg/utils/role"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initialize(t *testing.T) (context.Context, *UserService) {
	t.Helper()
	r := repos.NewUserMockRepo()
	return context.Background(), NewUserService(r)
}

func TestCreate(t *testing.T) {
	ctx, service := initialize(t)
	_, err := service.Create(ctx, "Joe", "jojo@gmail.com", "jojojo", role.User)
	assert.NoError(t, err)

	u, err := service.GetByName(ctx, "Joe")
	assert.NoError(t, err)
	assert.Equal(t, "jojo@gmail.com", u.Email)
}

func TestUpdateName(t *testing.T) {
	ctx, service := initialize(t)
	_, err := service.Create(ctx, "Joe", "jojo@gmail.com", "jojojo", role.User)
	assert.NoError(t, err)

	err = service.UpdateName(ctx, 1, "kkk")
	assert.NoError(t, err)

	u, err := service.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "kkk", u.Name)
}

func TestUpdateEmail(t *testing.T) {
	ctx, service := initialize(t)
	_, err := service.Create(ctx, "Joe", "jojo@gmail.com", "jojojo", role.User)
	assert.NoError(t, err)

	err = service.UpdateEmail(ctx, 1, "kkk")
	assert.NoError(t, err)

	u, err := service.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "kkk", u.Email)
}

func TestList(t *testing.T) {
	ctx, service := initialize(t)
	_, err := service.Create(ctx, "Joe", "jojo@gmail.com", "jojojo", role.User)
	assert.NoError(t, err)

	users, err := service.GetList(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
}

func TestDelete(t *testing.T) {
	ctx, service := initialize(t)
	_, err := service.Create(ctx, "Joe", "jojo@gmail.com", "jojojo", role.User)
	assert.NoError(t, err)

	err = service.Delete(ctx, 1)
	assert.NoError(t, err)

	users, err := service.GetList(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(users))
}
