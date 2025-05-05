package health

import (
	"context"
)

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Ping(ctx context.Context) error {
	return nil
}
