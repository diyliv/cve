package usecase

import "github.com/diyliv/cve/internal/cve"

type cveUC struct {
	postgresRepo cve.PostgresUsecase
}

func NewCveUC(postgresRepo cve.PostgresRepository) *cveUC {
	return &cveUC{
		postgresRepo: postgresRepo,
	}
}
