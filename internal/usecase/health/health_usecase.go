package health

import (
	"context"
	"fmt"

	"github.com/alif-academy-uz/auth-Cora23tt/internal/repository/health"
)

type Service struct {
	repo *health.Repo
}

func NewService(h *health.Repo) *Service {
	return &Service{repo: h}
}

func (s *Service) Check(ctx context.Context) error {
	return fmt.Errorf("health check failed: %w", s.repo.Ping(ctx))
}
