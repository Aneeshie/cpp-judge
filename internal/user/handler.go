package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{

		service: service,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req UserInput

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//TODO: better error handling

	userId, err := h.service.RegisterUser(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": userId})
}

func (h *Handler) SignIn(c *gin.Context) {
	var req UserInput

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	passwordMatch := h.service.Login(c.Request.Context(), req.Email, req.HashPassword)

	if !passwordMatch {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": req.Email})

}
