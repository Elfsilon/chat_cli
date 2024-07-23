package repos

import (
	"context"
	"server/internal/auth/gen/ent"
	"server/internal/auth/gen/ent/session"
	"time"
)

type SessionRepo struct {
	client *ent.Client
}

func NewSessionRepo(client *ent.Client) *SessionRepo {
	return &SessionRepo{client}
}

func (r *SessionRepo) Create(ctx context.Context, userID int, refreshToken string, expiresAt time.Time) (int, error) {
	s, err := r.client.Session.
		Create().
		SetUserID(userID).
		SetRefreshToken(refreshToken).
		SetExpiredAt(expiresAt).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return s.ID, nil
}

func (r *SessionRepo) GetByID(ctx context.Context, id int) (*ent.Session, error) {
	return r.client.Session.Query().Where(session.ID(id)).Only(ctx)
}

func (r *SessionRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*ent.Session, error) {
	return r.client.Session.Query().
		Where(
			session.RefreshToken(refreshToken),
		).
		Only(ctx)
}

func (r *SessionRepo) UpdateRefreshToken(ctx context.Context, id int, refreshToken string, expiredAt time.Time) error {
	return r.client.Session.
		UpdateOneID(id).
		SetRefreshToken(refreshToken).
		SetExpiredAt(expiredAt).
		Exec(ctx)
}
