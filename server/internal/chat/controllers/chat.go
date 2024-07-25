package controllers

import (
	"context"
	"errors"
	"fmt"
	"server/internal/chat/gen/chat"
	"server/internal/chat/services"
	ctxutil "server/pkg/utils/context_utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrConnectionFailedNotAMember = errors.New("unable to connect to the chat - you are not a member ")
	ErrSendingFailedNotAMember    = errors.New("unable to send a message to the chat - you are not a member")
)

type ChatController struct {
	chatService *services.ChatService
	chat.UnimplementedChatServer
}

func NewChatController(cs *services.ChatService) *ChatController {
	return &ChatController{
		chatService: cs,
	}
}

func (c *ChatController) HealthCheck(ctx context.Context, req *chat.HealthRequest) (*chat.HealthResponse, error) {
	return &chat.HealthResponse{}, nil
}

func (s *ChatController) Create(ctx context.Context, req *chat.CreateRequest) (*chat.CreateResponse, error) {
	ids64 := req.GetUserid()
	ids := make([]int, len(ids64))
	for _, id := range ids64 {
		ids = append(ids, int(id))
	}

	id, err := s.chatService.Create(ctx, req.GetName(), ids)
	if err != nil {
		return nil, err
	}

	response := &chat.CreateResponse{
		Id: int64(id),
	}

	return response, nil
}

func (s *ChatController) List(ctx context.Context, req *chat.ListRequest) (*chat.ListResponse, error) {
	userID, err := ctxutil.GetValue[int64](ctx, ctxutil.UserID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ls, err := s.chatService.List(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &chat.ListResponse{Chats: ls}
	return res, nil
}

func (s *ChatController) Conect(req *chat.ConnectRequest, stream chat.Chat_ConectServer) error {
	userID, err := ctxutil.GetValue[int64](stream.Context(), ctxutil.UserID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	userID, chatID := userID, int(req.GetChatID())

	isMember, err := s.chatService.IsMember(stream.Context(), chatID, int(userID))
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if !isMember {
		return status.Error(codes.PermissionDenied, ErrConnectionFailedNotAMember.Error())
	}

	return s.chatService.Connect(stream.Context(), chatID, int(userID), stream)
}

func (s *ChatController) SendMessage(ctx context.Context, req *chat.SendRequest) (*chat.SendResponse, error) {
	userID, err := ctxutil.GetValue[int64](ctx, ctxutil.UserID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	m := req.GetMessage()
	chatID := int(req.GetChatID())

	isMember, err := s.chatService.IsMember(ctx, chatID, int(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !isMember {
		return nil, status.Error(codes.PermissionDenied, ErrSendingFailedNotAMember.Error())
	}

	fmt.Printf("Send called, code = %v\n", m.GetColor())

	err = s.chatService.SendMessage(ctx, chatID, int(userID), m.GetUserName(), m.GetText(), m.GetColor())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &chat.SendResponse{}, nil
}

func (s *ChatController) Delete(ctx context.Context, req *chat.DeleteRequest) (*chat.DeleteResponse, error) {
	if err := s.chatService.Delete(ctx, int(req.GetChatID())); err != nil {
		return nil, err
	}

	return &chat.DeleteResponse{}, nil
}
