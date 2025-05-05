package health

import (
	"context"

	"github.com/alif-academy-uz/auth-Cora23tt/internal/repository/health"
)

type Service struct {
	repo *health.Repo
}

func NewService(h *health.Repo) *Service {
	return &Service{repo: h}
}

func (s *Service) Check(ctx context.Context) error {
	return s.repo.Ping(ctx)
}
