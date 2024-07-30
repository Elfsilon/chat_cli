package services

import (
	"context"
	"encoding/json"
	"fmt"
	"server/internal/chat/broker"
	"server/internal/chat/gen/chat"
	"server/internal/chat/repos"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatID int
type UserID int
type MessageStream chat.Chat_ConectServer
type UserMap map[UserID]MessageStream
type ChatMap map[ChatID]UserMap

const chatMessagesTopic = "chat_messages"

type TopicEvent struct {
	Type      int32     `json:"type"`
	ChatID    int       `json:"chat_id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"name"`
	Text      string    `json:"text"`
	Color     int32     `json:"color"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatService struct {
	broker   broker.MessageBroker
	chatRepo repos.Chat
	chatsMap ChatMap
	mu       sync.Mutex
}

func NewChatService(b broker.MessageBroker, cr repos.Chat) *ChatService {
	return &ChatService{
		broker:   b,
		chatRepo: cr,
		chatsMap: make(ChatMap),
		mu:       sync.Mutex{},
	}
}

func (s *ChatService) Init() (func() error, error) {
	ch, unsubscribe, err := s.broker.Subscribe(chatMessagesTopic)
	if err != nil {
		return nil, err
	}

	go s.listenMessages(ch)

	return unsubscribe, nil
}

func (s *ChatService) listenMessages(ch <-chan []byte) {
	for msg := range ch {
		var tm TopicEvent
		if err := json.Unmarshal(msg, &tm); err != nil {
			fmt.Printf("failed to unmarshal received message: %v\n", err)
			continue
		}

		chatID := ChatID(tm.ChatID)
		message := &chat.ChatEvent{
			UserID:    int64(tm.UserID),
			UserName:  tm.UserName,
			Text:      tm.Text,
			Color:     tm.Color,
			Type:      chat.EventType(tm.Type),
			Timestamp: timestamppb.New(tm.Timestamp),
		}

		s.mu.Lock()
		if c, ok := s.chatsMap[chatID]; ok {
			for _, stream := range c {
				stream.Send(message)
			}
		}
		s.mu.Unlock()
	}
}

func (s *ChatService) Create(ctx context.Context, name string, users []int) (int, error) {
	return s.chatRepo.Create(ctx, name, users)
}

func (s *ChatService) List(ctx context.Context, userID int64) ([]*chat.ChatInfo, error) {
	ls, err := s.chatRepo.GetList(ctx, int(userID))
	if err != nil {
		return nil, err
	}

	chats := make([]*chat.ChatInfo, len(ls))
	for i, ch := range ls {
		chats[i] = &chat.ChatInfo{
			Id:    int64(ch.ID),
			Title: ch.Name,
		}
	}

	return chats, nil
}

func (s *ChatService) IsMember(ctx context.Context, chatID, userID int) (bool, error) {
	chat, err := s.chatRepo.Get(ctx, chatID)
	if err != nil {
		return false, err
	}

	for _, id := range chat.Users {
		if userID == id {
			return true, nil
		}
	}

	return false, nil
}

func (s *ChatService) Connect(ctx context.Context, chatID, userID int, stream MessageStream) error {
	cid, uid := ChatID(chatID), UserID(userID)

	if err := s.connectUser(cid, uid, stream); err != nil {
		return err
	}
	s.sendEvent(int(chatID), userID, chat.EventType_Info, fmt.Sprintf("user %v has been connected", userID))
	defer func() {
		s.disconnectUser(cid, uid)
		s.sendEvent(int(chatID), userID, chat.EventType_Info, fmt.Sprintf("user %v has been disconnected", userID))
	}()

	<-stream.Context().Done()

	return nil
}

func (s *ChatService) connectUser(chatID ChatID, userID UserID, stream MessageStream) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.chatsMap[chatID]; !ok {
		s.chatsMap[chatID] = make(UserMap)
	}

	if _, ok := s.chatsMap[chatID][userID]; ok {
		return fmt.Errorf("user %v is already connected to the chat %v", userID, chatID)
	}

	s.chatsMap[chatID][userID] = stream
	fmt.Printf("user %v has been connected to the chat %v\n", userID, chatID)
	return nil
}

func (s *ChatService) disconnectUser(chatID ChatID, userID UserID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.chatsMap[chatID]; !ok {
		return
	}
	delete(s.chatsMap[chatID], userID)
	fmt.Printf("user %v has been deleted from the chat %v\n", userID, chatID)
}

func (s *ChatService) SendMessage(ctx context.Context, chatID int, userID int, username, text string, color int32) error {
	err := s.chatRepo.AddMessage(ctx, chatID, userID, username, text)
	if err != nil {
		return err
	}

	return s.publish(TopicEvent{
		ChatID:    chatID,
		UserID:    userID,
		UserName:  username,
		Text:      text,
		Color:     color,
		Type:      int32(chat.EventType_Message),
		Timestamp: time.Now(),
	})
}

func (s *ChatService) sendEvent(chatID, userID int, t chat.EventType, text string) error {
	return s.publish(TopicEvent{
		ChatID:    chatID,
		UserID:    userID,
		UserName:  "",
		Text:      text,
		Color:     0,
		Type:      int32(t),
		Timestamp: time.Now(),
	})
}

func (s *ChatService) publish(event TopicEvent) error {
	msgBytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}
	return s.broker.Publish(chatMessagesTopic, msgBytes)
}

func (s *ChatService) Delete(ctx context.Context, id int) error {
	return s.chatRepo.Delete(ctx, id)
}
