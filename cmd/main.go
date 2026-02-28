package main

import (
	"cppjudge/internal/config"
	"cppjudge/internal/server"
)

func main() {
	r := server.NewServer()
	cfg := config.Load()

	r.Run(":" + cfg.Port)
}
