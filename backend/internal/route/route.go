package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jirawatchn/resume-backend/internal/resume"
)

func NewRouter(h *resume.Handler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/resumes", h.Upload) // ผูก Route เข้ากับ Controller
	}

	return r
}
