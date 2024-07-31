package repos

import (
	"context"
	"server/internal/user/gen/ent"
	"server/pkg/utils/role"
)

type UserMockRepo struct {
	lastID int
	users  map[int]*ent.User
}

func NewUserMockRepo() *UserMockRepo {
	return &UserMockRepo{
		users: make(map[int]*ent.User, 0),
	}
}

func (r *UserMockRepo) Create(ctx context.Context, name string, email string, passwordHash string, role role.Role) (int, error) {
	r.lastID++
	u := &ent.User{
		ID:       r.lastID,
		Name:     name,
		Email:    email,
		Password: passwordHash,
		Role:     role,
	}
	r.users[r.lastID] = u
	return u.ID, nil
}

func (r *UserMockRepo) GetByID(ctx context.Context, id int) (*ent.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, &ent.NotFoundError{}
	}
	return u, nil
}

func (r *UserMockRepo) GetByName(ctx context.Context, name string) (*ent.User, error) {
	for _, u := range r.users {
		if u.Name == name {
			return u, nil
		}
	}
	return nil, &ent.NotFoundError{}
}

func (r *UserMockRepo) GetList(ctx context.Context) ([]*ent.User, error) {
	users := make([]*ent.User, 0)
	for _, u := range r.users {
		users = append(users, u)
	}
	return users, nil
}

func (r *UserMockRepo) UpdateName(ctx context.Context, id int, name string) error {
	r.users[id].Name = name
	return nil
}

func (r *UserMockRepo) UpdateEmail(ctx context.Context, id int, email string) error {
	r.users[id].Email = email
	return nil
}

func (r *UserMockRepo) Delete(ctx context.Context, id int) error {
	delete(r.users, id)
	return nil
}
