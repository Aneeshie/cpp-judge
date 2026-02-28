package server

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	r.GET("/health", GetHealth)

	return r
}
