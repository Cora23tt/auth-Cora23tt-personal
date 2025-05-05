package rest

import (
	"net/http"

	"github.com/alif-academy-uz/auth-Cora23tt/internal/rest/handlers/health"
	"github.com/gin-gonic/gin"
)

func NewServer(mux *gin.Engine, h *health.Handler) *Server {
	return &Server{
		mux: mux,
		h:   h,
	}
}

type Server struct {
	mux *gin.Engine
	h   *health.Handler
}

func (s *Server) Init() {

	const baseUrl = "/api/v1"

	s.mux.Use(gin.Recovery())
	s.mux.Use(gin.Logger())

	publicGroup := s.mux.Group(baseUrl + "/public")
	{
		publicGroup.GET("/health", s.h.HealthCheck)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
