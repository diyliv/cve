package cve

import (
	"context"

	"github.com/diyliv/cve/internal/models"
)

type PostgresUsecase interface {
	Register(ctx context.Context, creds models.Creds) error
	Login(ctx context.Context, creds models.Creds) (string, error)
}
