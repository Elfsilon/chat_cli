package utils

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/metadata"
)

type TokenUpdater struct {
	token   string
	ttl     time.Duration
	manager *TokenManager
}

func NewTokenUpdater(tm *TokenManager, ttl time.Duration) *TokenUpdater {
	return &TokenUpdater{manager: tm, ttl: ttl}
}

func (t *TokenUpdater) Token() string {
	return t.token
}

func (t *TokenUpdater) Context(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "s-authorization", fmt.Sprintf("Service %v", t.token))
}

func (t *TokenUpdater) UpdateToken() error {
	token, err := t.manager.GenerateServiceToken()
	if err != nil {
		return fmt.Errorf("unable to generate service token: %v", err)
	}
	t.token = token
	return nil
}

func (t *TokenUpdater) Run() {
	ticker := time.NewTicker(t.ttl)
	for range ticker.C {
		if err := t.UpdateToken(); err != nil {
			fmt.Println(err)
		}
	}
}
