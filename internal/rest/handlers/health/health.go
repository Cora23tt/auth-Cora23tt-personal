package health

import (
	"log"
	"net/http"

	"github.com/alif-academy-uz/auth-Cora23tt/internal/usecase/health"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *health.Service
}

func NewHandler(service *health.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	if err := h.service.Check(c.Request.Context()); err != nil {
		http.Error(c.Writer, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	if _, err := c.Writer.WriteString("Ping Pong"); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
