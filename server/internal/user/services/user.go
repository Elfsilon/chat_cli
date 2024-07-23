package services

import (
	"context"
	"server/internal/user/gen/ent"
	"server/internal/user/repos"
	"server/pkg/utils/hasher"
	"server/pkg/utils/role"
)

type UserService struct {
	userRepo *repos.UserRepo
}

func NewUserService(ur *repos.UserRepo) *UserService {
	return &UserService{
		userRepo: ur,
	}
}

func (s *UserService) Create(ctx context.Context, name string, email string, password string, role role.Role) (int, error) {
	phash, err := hasher.HashPassword(password)
	if err != nil {
		return -1, err
	}

	return s.userRepo.Create(ctx, name, email, phash, role)
}

func (r *UserService) GetByID(ctx context.Context, id int) (*ent.User, error) {
	return r.userRepo.GetByID(ctx, id)
}

func (r *UserService) GetByName(ctx context.Context, name string) (*ent.User, error) {
	return r.userRepo.GetByName(ctx, name)
}

func (r *UserService) GetList(ctx context.Context) ([]*ent.User, error) {
	return r.userRepo.GetList(ctx)
}

func (r *UserService) UpdateName(ctx context.Context, id int, name string) error {
	return r.userRepo.UpdateName(ctx, id, name)
}

func (r *UserService) UpdateEmail(ctx context.Context, id int, email string) error {
	return r.userRepo.UpdateEmail(ctx, id, email)
}

func (r *UserService) Delete(ctx context.Context, id int) error {
	return r.userRepo.Delete(ctx, id)
}
