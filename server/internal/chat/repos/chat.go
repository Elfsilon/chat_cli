package repos

import (
	"context"
	"server/internal/chat/gen/ent"
	"server/internal/chat/gen/ent/chat"
	"server/internal/chat/gen/ent/chatmessage"
)

type ChatRepo struct {
	client *ent.Client
}

func NewChatRepo(client *ent.Client) *ChatRepo {
	return &ChatRepo{client}
}

func (r *ChatRepo) Create(ctx context.Context, name string, users []int) (int, error) {
	s, err := r.client.Chat.
		Create().
		SetName(name).
		SetUsers(users).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return s.ID, nil
}

func (r *ChatRepo) AddMessage(ctx context.Context, chatID, userID int, username, message string) error {
	_, err := r.client.ChatMessage.
		Create().
		SetOwnerID(chatID).
		SetUserid(userID).
		SetUsername(username).
		SetMessage(message).
		Save(ctx)

	return err
}

func (r *ChatRepo) Get(ctx context.Context, chatID int) (*ent.Chat, error) {
	return r.client.Chat.Query().Where(chat.ID(chatID)).Only(ctx)
}

func (r *ChatRepo) GetHistory(ctx context.Context, chatID int) ([]*ent.ChatMessage, error) {
	history, err := r.client.ChatMessage.
		Query().
		Where(chatmessage.HasOwnerWith(chat.ID(chatID))).
		Order(chatmessage.ByTime()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (r *ChatRepo) Delete(ctx context.Context, id int) error {
	_, err := r.client.ChatMessage.
		Delete().
		Where(chatmessage.HasOwnerWith(chat.ID(id))).
		Exec(ctx)

	if err != nil {
		return err
	}

	return r.client.Chat.DeleteOneID(id).Exec(ctx)
}
