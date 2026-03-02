package submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) MakeSubmission(c *gin.Context) {
	var req SubmissionInput

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserID == "" || req.ProblemId == "" || req.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid inputs"})
		return
	}

	created, err := h.repo.MakeSubmission(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not make the submission"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"submission_id": created})

}
