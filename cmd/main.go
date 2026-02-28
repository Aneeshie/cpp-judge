package main

import (
	"cppjudge/internal/config"
	"cppjudge/internal/database"
	"cppjudge/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	conn, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	log.Println("database connected successfully")

	r := server.NewServer(conn)

	r.Run(":" + cfg.Port)
}
