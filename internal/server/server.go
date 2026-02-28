package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewServer(conn *pgx.Conn) *gin.Engine {
	r := gin.Default()

	r.GET("/health", GetHealth)

	return r
}
