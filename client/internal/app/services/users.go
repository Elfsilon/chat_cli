package services

import (
	"chat_cli/internal/app/gen/user"
	"chat_cli/internal/app/models"
	"context"
)

type UserService struct {
	client user.UserServiceClient
}

func NewUserService(client user.UserServiceClient) *UserService {
	return &UserService{client}
}

func (s *UserService) Create(ctx context.Context, name, email, password string, role int32) (int64, error) {
	req := &user.CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     user.Role(role),
	}
	res, err := s.client.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (s *UserService) getByReq(ctx context.Context, req *user.GetRequest) (*models.User, error) {
	res, err := s.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		ID:    res.GetId(),
		Name:  res.GetName(),
		Email: res.GetEmail(),
		Role:  int32(res.GetRole()),
	}
	return u, nil
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	req := &user.GetRequest{Id: id}
	return s.getByReq(ctx, req)
}

func (s *UserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	req := &user.GetRequest{Username: name}
	return s.getByReq(ctx, req)
}

func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	res, err := s.client.GetList(ctx, &user.GetListRequest{})
	if err != nil {
		return nil, err
	}

	ls := res.GetUsers()
	users := make([]models.User, len(ls))
	for i, u := range ls {
		users[i] = models.User{
			ID:    u.GetId(),
			Name:  u.GetName(),
			Email: u.GetEmail(),
			Role:  int32(u.GetRole()),
		}
	}

	return users, nil
}

func (s *UserService) Update(ctx context.Context, id int64, email, name string) error {
	req := &user.UpdateRequest{
		Id:    id,
		Name:  name,
		Email: email,
	}
	_, err := s.client.Update(ctx, req)
	return err
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	req := &user.DeleteRequest{Id: id}
	_, err := s.client.Delete(ctx, req)
	return err
}
