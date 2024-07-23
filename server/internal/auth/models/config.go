package models

import "time"

type Config struct {
	Db     DatabaseConfig
	Auth   AuthConfig
	Server ServerConfig
}

type AuthConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	JWTSecret       string
	ServiceSecret   string
}

type DatabaseConfig struct {
	ConnURL string
}

type ServerConfig struct {
	Host string
	Port int
}
