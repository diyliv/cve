package interceptors

import (
	"context"

	"github.com/diyliv/cve/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type interceptor struct {
	logger *zap.Logger
	cfg    *config.Config
}

func NewInterceptor(logger *zap.Logger, cfg *config.Config) *interceptor {
	return &interceptor{
		logger: logger,
		cfg:    cfg,
	}
}

func (i *interceptor) Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	return nil, nil
}
