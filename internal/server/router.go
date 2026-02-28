package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"status": "Ok",
	})
}
