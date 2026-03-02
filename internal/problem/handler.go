package problem

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

func (h *Handler) CreateProblem(c *gin.Context) {
	var req Problem

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.repo.CreateProblem(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create problem"})
		return
	}

	c.JSON(http.StatusCreated, created)

}

func (h *Handler) GetProblems(c *gin.Context) {
	problems, err := h.repo.GetProblems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch problems",
		})
		return
	}

	c.JSON(http.StatusOK, problems)
}

func (h *Handler) MakeSubmission(c *gin.Context) {
	var req SubmissionInput

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserID == "" || req.ProblemId == "" || req.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid inputs"})
	}

	created, err := h.repo.MakeSubmission(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not make the submission"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"submission_id": created})

}
