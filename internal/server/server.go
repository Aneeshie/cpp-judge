package server

import (
	"cppjudge/internal/problem"
	"cppjudge/internal/submission"
	"cppjudge/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewServer(conn *pgx.Conn) *gin.Engine {
	r := gin.Default()

	problemRepo := problem.NewRepository(conn)
	problemHandler := problem.NewHandler(problemRepo)

	submissionRepo := submission.NewRepository(conn)
	submissionHandler := submission.NewHandler(submissionRepo)

	userRepo := user.NewRepository(conn)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	//problems

	//TODO: duplication problems
	r.POST("/problems", problemHandler.CreateProblem)

	//TODO: NOT ABLE TO FETCH PROBLEMS
	r.GET("/problems", problemHandler.GetProblems)

	//submissions
	r.POST("/submissions", submissionHandler.MakeSubmission)

	//user
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/signin", userHandler.SignIn)

	return r
}
