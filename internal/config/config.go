package config

import (
	"fmt"
	"os"
)

type ServerConfig struct {
	Port        string
	DatabaseURL string
}

func Load() *ServerConfig {
	// check the env for port, if not present hardcode
	// right now the ServerConfig only contains port, later add stuff if needed

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_port := os.Getenv("POSTGRES_PORT")
	postgres_database := os.Getenv("POSTGRES_DATABASE")
	postgres_host := os.Getenv("POSTGRES_HOST")

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgres_user, postgres_password, postgres_host, postgres_port, postgres_database)

	return &ServerConfig{
		Port:        port,
		DatabaseURL: db_url,
	}

}
