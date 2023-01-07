package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/diyliv/cve/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func (i *interceptor) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
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

func (i *interceptor) CheckToken(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.FullMethod == "/cve.CveService/Register" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Retrieving metadata in failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Go to register to get your personal token")
	}

	token := authHeader[0]
	fmt.Println(token)

	reply, err := handler(ctx, req)
	return reply, err
}
