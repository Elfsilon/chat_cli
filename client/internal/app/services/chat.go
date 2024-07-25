package services

import (
	"chat_cli/internal/app/gen/chat"
	"chat_cli/internal/app/models"
	"context"
	"fmt"
	"io"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatService struct {
	client chat.ChatClient
}

func NewChatService(client chat.ChatClient) *ChatService {
	return &ChatService{client}
}

func (s *ChatService) Create(ctx context.Context, title string, users []int64) (int64, error) {
	req := &chat.CreateRequest{Name: title, Userid: users}
	res, err := s.client.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (s *ChatService) Connect(ctx context.Context, chatID int64) (<-chan models.ChatMessage, error) {
	req := &chat.ConnectRequest{ChatID: chatID}
	stream, err := s.client.Conect(ctx, req)
	if err != nil {
		return nil, err
	}

	ch := make(chan models.ChatMessage)
	go func() {
		for {
			if ctx.Err() != nil || stream.Context().Err() != nil {
				fmt.Printf("context error: %v\n", err)
				break
			}

			m, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("EOF error: %v\n", err)
				break
			}
			if err != nil {
				fmt.Printf("reading the stream message error: %v\n", err)
				break
			}

			ch <- models.ChatMessage{
				Author: m.GetUserName(),
				Text:   m.GetText(),
			}
		}
		close(ch)
	}()

	return ch, nil
}

func (s *ChatService) List(ctx context.Context) ([]models.ChatInfo, error) {
	res, err := s.client.List(ctx, &chat.ListRequest{})
	if err != nil {
		return nil, err
	}

	chats := res.GetChats()
	info := make([]models.ChatInfo, len(chats))
	for i, chat := range chats {
		info[i] = models.ChatInfo{
			ID:    chat.Id,
			Title: chat.Title,
		}
	}

	return info, nil
}

func (s *ChatService) Delete(ctx context.Context, id int64) error {
	req := &chat.DeleteRequest{ChatID: id}
	_, err := s.client.Delete(ctx, req)
	return err
}

func (s *ChatService) Send(ctx context.Context, chatID int64, message models.ChatMessage) error {
	req := &chat.SendRequest{
		ChatID: chatID,
		Message: &chat.ChatMessage{
			UserName:  message.Author,
			Text:      message.Text,
			Timestamp: timestamppb.Now(),
		},
	}
	_, err := s.client.SendMessage(ctx, req)
	return err
}
