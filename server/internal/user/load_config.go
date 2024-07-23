package app

import (
	"server/internal/user/models"
	"server/pkg/utils/env"
)

func (a *App) LoadConfig() {
	dbConnURL := env.String("DB_CONN_URL")

	host := env.OptionalString("HOST")
	port := env.Int("PORT")

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
			ServiceSecret: serviceSecret,
		},
	}
}
