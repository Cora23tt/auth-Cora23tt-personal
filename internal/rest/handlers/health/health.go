package health

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	if _, err := c.Writer.WriteString("Ping Pong"); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
