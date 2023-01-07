package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/diyliv/cve/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	start := time.Now()

	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	i.logger.Info(fmt.Sprintf("Method: %s, Time: %v, Metadata: %v, Err: %v\n",
		info.FullMethod,
		time.Since(start),
		md,
		err))
	return reply, err
}
