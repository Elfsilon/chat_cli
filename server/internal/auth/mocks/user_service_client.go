package mocks

import (
	"context"
	"server/internal/auth/gen/user"

	"google.golang.org/grpc"
)

type UserServiceMockClient struct {
	lastID int
	users  map[int]*user.UserModel
}

func NewUserServiceMockClient() *UserServiceMockClient {
	return &UserServiceMockClient{
		users: make(map[int]*user.UserModel, 0),
	}
}

func (a *UserServiceMockClient) HealthCheck(ctx context.Context, in *user.HealthRequest, opts ...grpc.CallOption) (*user.HealthResponse, error) {
	return &user.HealthResponse{}, nil
}

func (a *UserServiceMockClient) Create(ctx context.Context, in *user.CreateUserRequest, opts ...grpc.CallOption) (*user.CreateUserResponse, error) {
	a.lastID++
	u := &user.UserModel{
		Id:           int64(a.lastID),
		Name:         in.GetName(),
		Email:        in.GetEmail(),
		Role:         in.Role,
		PasswordHash: in.GetPassword(),
	}
	a.users[a.lastID] = u

	return &user.CreateUserResponse{Id: u.Id}, nil
}

func (a *UserServiceMockClient) Get(ctx context.Context, in *user.GetRequest, opts ...grpc.CallOption) (*user.UserModel, error) {
	if in.GetUsername() != "" {
		for _, u := range a.users {
			if u.Name == in.GetUsername() {
				return u, nil
			}
		}
	}
	return a.users[int(in.GetId())], nil
}

func (a *UserServiceMockClient) GetList(ctx context.Context, in *user.GetListRequest, opts ...grpc.CallOption) (*user.GetListResponse, error) {
	return &user.GetListResponse{}, nil
}

func (a *UserServiceMockClient) Update(ctx context.Context, in *user.UpdateRequest, opts ...grpc.CallOption) (*user.UpdateResponse, error) {
	return &user.UpdateResponse{}, nil
}

func (a *UserServiceMockClient) Delete(ctx context.Context, in *user.DeleteRequest, opts ...grpc.CallOption) (*user.DeleteResponse, error) {
	return &user.DeleteResponse{}, nil
}
