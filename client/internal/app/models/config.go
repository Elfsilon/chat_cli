package models

import "time"

type Config struct {
	AuthServiceUrl string
	ChatServiceUrl string
	UserServiceUrl string

	AccessTokenTTL time.Duration
}
