package app

import (
	"server/internal/chat/models"
	"server/pkg/utils/env"
)

func (a *App) LoadConfig() {
	dbConnURL := env.String("DB_CONN_URL")
	natsURL := env.String("NATS_URL")

	host := env.OptionalString("HOST")
	port := env.Int("PORT")

	serviceSecret := env.String("SERVICE_SECRET")

	a.cfg = models.Config{
		NatsUrl: natsURL,
		Server: models.ServerConfig{
			Host: host,
			Port: port,
		},
		Db: models.DatabaseConfig{
			ConnURL: dbConnURL,
		},
		Auth: models.AuthConfig{
			ServiceSecret: serviceSecret,
		},
	}
}
