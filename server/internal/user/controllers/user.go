package controllers

import (
	"context"
	"errors"
	"server/internal/user/gen/ent"
	"server/internal/user/gen/user"
	"server/internal/user/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrUserIdOrNameMissed = errors.New("either user_id or user_name has to be passed")

type UserController struct {
	userService *services.UserService
	user.UnimplementedUserServiceServer
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (c *UserController) HealthCheck(ctx context.Context, req *user.HealthRequest) (*user.HealthResponse, error) {
	return &user.HealthResponse{}, nil
}

func (c *UserController) Create(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	userID, err := c.userService.Create(ctx, req.GetName(), req.GetEmail(), req.GetPassword(), int(req.GetRole()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response := &user.CreateUserResponse{
		Id: int64(userID),
	}

	return response, nil
}

// TODO
func (c *UserController) Get(ctx context.Context, req *user.GetRequest) (*user.UserModel, error) {
	userID, userName := int(req.GetId()), req.GetUsername()

	var u *ent.User
	var err error

	if userID != 0 {
		u, err = c.userService.GetByID(ctx, userID)
	} else if userName != "" {
		u, err = c.userService.GetByName(ctx, userName)
	} else {
		u, err = nil, ErrUserIdOrNameMissed
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response := &user.UserModel{
		Id:           int64(u.ID),
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.Password,
		Role:         user.Role(u.Role),
		CreatedAt:    timestamppb.New(u.CreatedAt),
		UpdatedAt:    timestamppb.New(u.UpdatedAt),
	}

	return response, nil
}

func (c *UserController) GetList(ctx context.Context, req *user.GetListRequest) (*user.GetListResponse, error) {
	u, err := c.userService.GetList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	users := make([]*user.UserModel, 0)
	for _, uent := range u {
		user := &user.UserModel{
			Id:        int64(uent.ID),
			Name:      uent.Name,
			Email:     uent.Email,
			Role:      user.Role(uent.Role),
			CreatedAt: timestamppb.New(uent.CreatedAt),
			UpdatedAt: timestamppb.New(uent.UpdatedAt),
		}
		users = append(users, user)
	}

	response := &user.GetListResponse{
		Users: users,
	}

	return response, nil
}

func (c *UserController) Update(ctx context.Context, req *user.UpdateRequest) (*user.UpdateResponse, error) {
	if req.Name != "" {
		if err := c.userService.UpdateName(ctx, int(req.GetId()), req.GetName()); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if req.Email != "" {
		if err := c.userService.UpdateEmail(ctx, int(req.GetId()), req.GetEmail()); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &user.UpdateResponse{}, nil
}

func (c *UserController) Delete(ctx context.Context, req *user.DeleteRequest) (*user.DeleteResponse, error) {
	if err := c.userService.Delete(ctx, int(req.Id)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &user.DeleteResponse{}, nil
}
