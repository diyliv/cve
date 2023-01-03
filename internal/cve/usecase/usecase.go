package usecase

import (
	"context"

	"github.com/diyliv/cve/internal/cve"
	"github.com/diyliv/cve/internal/models"
)

type cveUC struct {
	postgresRepo cve.PostgresUsecase
}

func NewCveUC(postgresRepo cve.PostgresRepository) *cveUC {
	return &cveUC{
		postgresRepo: postgresRepo,
	}
}

func (cv *cveUC) Register(ctx context.Context, creds models.Creds) error {
	return cv.postgresRepo.Register(ctx, creds)
}

func (cv *cveUC) Login(ctx context.Context, creds models.Creds) (string, error) {
	return cv.postgresRepo.Login(ctx, creds)
}
