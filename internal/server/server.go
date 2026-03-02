package server

import (
	"cppjudge/internal/problem"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewServer(conn *pgx.Conn) *gin.Engine {
	r := gin.Default()

	problemRepo := problem.NewRepository(conn)
	problemHandler := problem.NewHandler(problemRepo)

	r.POST("/problems", problemHandler.CreateProblem)

	return r
}
