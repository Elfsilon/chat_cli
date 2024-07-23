package app

import (
	"server/internal/auth/models"
	"server/pkg/utils/env"
	"time"
)

func (a *App) LoadConfig() {
	dbConnURL := env.String("DB_CONN_URL")

	host := env.OptionalString("HOST")
	port := env.Int("PORT")

	accessTTL := env.Int("ACCESS_TOKEN_TTL")
	refreshTTL := env.Int("REFRESH_TOKEN_TTL")
	secret := env.String("JWT_SECRET")
	serviceSecret := env.String("SERVICE_SECRET")

	a.cfg = models.Config{
		Server: models.ServerConfig{
			Host: host,
			Port: port,
		},
		Db: models.DatabaseConfig{
			ConnURL: dbConnURL,
		},
		Auth: models.AuthConfig{
			AccessTokenTTL:  time.Duration(accessTTL) * time.Second,
			RefreshTokenTTL: time.Duration(refreshTTL) * time.Second,
			JWTSecret:       secret,
			ServiceSecret:   serviceSecret,
		},
	}
}
