package server

import (
	"cppjudge/internal/problem"
	"cppjudge/internal/submission"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewServer(conn *pgx.Conn) *gin.Engine {
	r := gin.Default()

	problemRepo := problem.NewRepository(conn)
	problemHandler := problem.NewHandler(problemRepo)

	submissionRepo := submission.NewRepository(conn)
	submissionHandler := submission.NewHandler(submissionRepo)

	r.POST("/problems", problemHandler.CreateProblem)
	r.POST("/submissions", submissionHandler.MakeSubmission)
	r.GET("/problems", problemHandler.GetProblems)

	return r
}
