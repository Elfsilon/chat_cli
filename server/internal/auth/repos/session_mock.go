package repos

import (
	"context"
	"server/internal/auth/gen/ent"
	"time"
)

type SessionMockRepo struct {
	lastID  int
	storage map[int]*ent.Session
}

func NewSessionMockRepo() *SessionMockRepo {
	return &SessionMockRepo{
		storage: make(map[int]*ent.Session, 0),
	}
}

func (r *SessionMockRepo) Create(ctx context.Context, userID int, refreshToken string, expiresAt time.Time) (int, error) {
	r.lastID++
	session := &ent.Session{
		ID:           r.lastID,
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiredAt:    expiresAt,
		CreatedAt:    time.Now(),
	}
	r.storage[r.lastID] = session
	return session.ID, nil
}

func (r *SessionMockRepo) GetByID(ctx context.Context, id int) (*ent.Session, error) {
	return r.storage[id], nil
}

func (r *SessionMockRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*ent.Session, error) {
	for _, s := range r.storage {
		if s.RefreshToken == refreshToken {
			return s, nil
		}
	}
	return nil, &ent.NotFoundError{}
}

func (r *SessionMockRepo) UpdateRefreshToken(ctx context.Context, id int, refreshToken string, expiredAt time.Time) error {
	r.storage[id].RefreshToken = refreshToken
	r.storage[id].ExpiredAt = expiredAt
	return nil
}
