package app

import (
	"chat_cli/internal/app/models"
	"chat_cli/pkg/env"
	"time"

	"github.com/joho/godotenv"
)

func (a *App) LoadConfig() {
	godotenv.Load("/Users/user/Desktop/Test/client/config/dev.env")

	authServiceURL := env.String("AUTH_URL")
	chatServiceURL := env.String("CHAT_URL")
	userServiceURL := env.String("USER_URL")

	accessTokenTTL := env.Int("ACCESS_TOKEN_TTL")

	jwtSecret := env.String("JWT_SECRET")

	a.config = models.Config{
		AuthServiceUrl: authServiceURL,
		ChatServiceUrl: chatServiceURL,
		UserServiceUrl: userServiceURL,
		AccessTokenTTL: time.Duration(accessTokenTTL) * time.Second,
		JwtSecret:      []byte(jwtSecret),
	}
}
