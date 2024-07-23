package repos

import (
	"context"
	"server/internal/user/gen/ent"
	"server/internal/user/gen/ent/user"
	"server/pkg/utils/role"
)

type UserRepo struct {
	client *ent.Client
}

func NewUserRepo(client *ent.Client) *UserRepo {
	return &UserRepo{client}
}

func (r *UserRepo) Create(ctx context.Context, name string, email string, passwordHash string, role role.Role) (int, error) {
	u, err := r.client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(passwordHash).
		SetRole(role).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return u.ID, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Query().Where(user.ID(id)).Only(ctx)
}

func (r *UserRepo) GetByName(ctx context.Context, name string) (*ent.User, error) {
	return r.client.User.Query().Where(user.Name(name)).Only(ctx)
}

func (r *UserRepo) GetList(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *UserRepo) UpdateName(ctx context.Context, id int, name string) error {
	return r.client.User.UpdateOneID(id).SetName(name).Exec(ctx)
}

func (r *UserRepo) UpdateEmail(ctx context.Context, id int, email string) error {
	return r.client.User.UpdateOneID(id).SetEmail(email).Exec(ctx)
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}
