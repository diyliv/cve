package grpc

import (
	"github.com/diyliv/cve/config"
	"github.com/diyliv/cve/internal/cve"
	"go.uber.org/zap"
)

type grpcservice struct {
	logger *zap.Logger
	cfg    *config.Config
	cveUC  cve.PostgresUsecase
}

func NewgRPCService(logger *zap.Logger, cfg *config.Config, cveUC cve.PostgresUsecase) *grpcservice {
	return &grpcservice{
		logger: logger,
		cfg:    cfg,
		cveUC:  cveUC,
	}
}
