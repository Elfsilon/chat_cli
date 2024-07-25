package repos

import (
	"context"
	"server/internal/chat/gen/ent"
	"server/internal/chat/gen/ent/chat"
	"server/internal/chat/gen/ent/chatmessage"
	"server/internal/chat/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
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
		return 0, err
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
	return r.client.Chat.
		Query().
		Where(chat.ID(chatID)).
		Only(ctx)
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

// TODO: Find a better place
func HasUser(id int) predicate.Chat {
	return predicate.Chat(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(chat.FieldUsers, id))
	})
}

func (r *ChatRepo) GetList(ctx context.Context, userID int) ([]*ent.Chat, error) {
	return r.client.Chat.
		Query().
		Where(HasUser(userID)).
		All(ctx)
}

func (r *ChatRepo) Delete(ctx context.Context, id int) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.ChatMessage.
		Delete().
		Where(chatmessage.HasOwnerWith(chat.ID(id))).
		Exec(ctx)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Chat.DeleteOneID(id).Exec(ctx); err != nil {
		tx.Rollback()
	}

	return tx.Commit()
}
