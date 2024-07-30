package repos

import (
	"context"
	"server/internal/chat/gen/ent"
)

type mockMessage struct {
	id       int
	userid   int
	username string
	message  string
}

type mockChat struct {
	id       int
	name     string
	users    []int
	messages []mockMessage
}

type ChatMockRepo struct {
	lastID int
	chats  map[int]*mockChat
}

func NewChatMockRepo() *ChatMockRepo {
	return &ChatMockRepo{
		chats: make(map[int]*mockChat, 0),
	}
}

func (r *ChatMockRepo) Create(ctx context.Context, name string, users []int) (int, error) {
	r.lastID++
	c := &mockChat{
		id:       r.lastID,
		name:     name,
		users:    users,
		messages: make([]mockMessage, 0),
	}
	r.chats[r.lastID] = c
	return c.id, nil
}

func (r *ChatMockRepo) AddMessage(ctx context.Context, chatID, userID int, username, message string) error {
	m := mockMessage{
		id:       len(r.chats[chatID].messages),
		userid:   userID,
		username: username,
		message:  message,
	}
	r.chats[chatID].messages = append(r.chats[chatID].messages, m)
	return nil
}

func (r *ChatMockRepo) Get(ctx context.Context, chatID int) (*ent.Chat, error) {
	chat, ok := r.chats[chatID]
	if !ok {
		return nil, &ent.NotFoundError{}
	}

	res := &ent.Chat{
		ID:    chat.id,
		Name:  chat.name,
		Users: chat.users,
	}
	return res, nil
}

func (r *ChatMockRepo) GetHistory(ctx context.Context, chatID int) ([]*ent.ChatMessage, error) {
	return []*ent.ChatMessage{}, nil
}

func (r *ChatMockRepo) GetList(ctx context.Context, userID int) ([]*ent.Chat, error) {
	res := make([]*ent.Chat, 0)
	for _, c := range r.chats {
		for _, uid := range c.users {
			if uid == userID {
				res = append(res, &ent.Chat{
					ID:    c.id,
					Name:  c.name,
					Users: c.users,
				})
			}
		}
	}
	return res, nil
}

func (r *ChatMockRepo) Delete(ctx context.Context, id int) error {
	delete(r.chats, id)
	return nil
}
