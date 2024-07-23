package models

type Config struct {
	Db     DatabaseConfig
	Auth   AuthConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	ConnURL string
}

type ServerConfig struct {
	Host string
	Port int
}

type AuthConfig struct {
	ServiceSecret string
}
