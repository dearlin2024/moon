package repository

import (
	"context"

	"github.com/aide-family/moon/app/demo/internal/biz/bo"
)

// PingRepo is a Greater repo.
type PingRepo interface {
	Ping(ctx context.Context, g *bo.Ping) (*bo.Ping, error)
}
