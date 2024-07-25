package services

import (
	"context"
	"encoding/json"
	"fmt"
	"server/internal/chat/gen/chat"
	"server/internal/chat/repos"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatID int
type UserID int
type MessageStream chat.Chat_ConectServer
type UserMap map[UserID]MessageStream
type ChatMap map[ChatID]UserMap

const chatMessagesTopic = "chat_messages"

type TopicMessage struct {
	ChatID    int       `json:"chat_id"`
	UserName  string    `json:"name"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatService struct {
	nc       *nats.Conn
	chatRepo *repos.ChatRepo
	chatsMap ChatMap
	mu       sync.Mutex
}

func NewChatService(nc *nats.Conn, cr *repos.ChatRepo) *ChatService {
	return &ChatService{
		nc:       nc,
		chatRepo: cr,
		chatsMap: make(ChatMap),
		mu:       sync.Mutex{},
	}
}

func (s *ChatService) Init() (func() error, error) {
	ch := make(chan *nats.Msg)
	sub, err := s.nc.ChanSubscribe(chatMessagesTopic, ch)
	if err != nil {
		return nil, err
	}

	go s.listenMessages(ch)

	return sub.Unsubscribe, nil
}

func (s *ChatService) listenMessages(ch chan *nats.Msg) {
	for msg := range ch {
		var tm TopicMessage
		if err := json.Unmarshal(msg.Data, &tm); err != nil {
			fmt.Printf("failed to unmarshal received message: %v\n", err)
			continue
		}

		chatID := ChatID(tm.ChatID)
		message := &chat.ChatMessage{
			UserName:  tm.UserName,
			Text:      tm.Text,
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

	s.connectUser(cid, uid, stream)
	defer s.disconnectUser(cid, uid)

	<-stream.Context().Done()

	return nil
}

func (s *ChatService) connectUser(chatID ChatID, userID UserID, stream MessageStream) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.chatsMap[chatID]; !ok {
		s.chatsMap[chatID] = make(UserMap)
	}
	s.chatsMap[chatID][userID] = stream
	fmt.Printf("user %v has been connected to the chat %v\n", userID, chatID)
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

func (s *ChatService) SendMessage(ctx context.Context, chatID int, userID int, username, text string) error {
	err := s.chatRepo.AddMessage(ctx, chatID, userID, username, text)
	if err != nil {
		return err
	}

	msg := TopicMessage{
		ChatID:    chatID,
		UserName:  username,
		Text:      text,
		Timestamp: time.Now(),
	}

	msgBytes, err := json.Marshal(&msg)
	if err != nil {
		return err
	}

	return s.nc.Publish(chatMessagesTopic, msgBytes)
}

func (s *ChatService) Delete(ctx context.Context, id int) error {
	return s.chatRepo.Delete(ctx, id)
}
