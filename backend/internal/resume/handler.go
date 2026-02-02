package resume

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler { return &Handler{service: s} }

func (h *Handler) Upload(c *gin.Context) {
	var req Resume
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, _ := h.service.ScreenResume(req)
	c.JSON(http.StatusOK, result)
}
