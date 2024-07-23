package app

import (
	"chat_cli/internal/app/models"
	"chat_cli/pkg/env"
	"time"
)

func (a *App) LoadConfig() {
	authServiceURL := env.String("AUTH_URL")
	chatServiceURL := env.String("CHAT_URL")
	userServiceURL := env.String("USER_URL")

	accessTokenTTL := env.Int("ACCESS_TOKEN_TTL")

	a.config = models.Config{
		AuthServiceUrl: authServiceURL,
		ChatServiceUrl: chatServiceURL,
		UserServiceUrl: userServiceURL,
		AccessTokenTTL: time.Duration(accessTokenTTL) * time.Second,
	}
}
