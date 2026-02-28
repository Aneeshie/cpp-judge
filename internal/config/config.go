package config

import "os"

type ServerConfig struct {
	Port string
}

func Load() *ServerConfig {
	// check the env for port, if not present hardcode
	// right now the ServerConfig only contains port, later add stuff if needed

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return &ServerConfig{
		Port: port,
	}

}
