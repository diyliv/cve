package server

import (
	"database/sql"
	"net"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/diyliv/cve/config"
	grpcservice "github.com/diyliv/cve/internal/cve/delivery/grpc"
	"github.com/diyliv/cve/internal/cve/repository"
	"github.com/diyliv/cve/internal/cve/usecase"
	"github.com/diyliv/cve/internal/interceptors"
	cvepb "github.com/diyliv/cve/proto/cve"
)

type server struct {
	logger   *zap.Logger
	cfg      *config.Config
	psqlConn *sql.DB
}

func NewServer(logger *zap.Logger, cfg *config.Config, psqlConn *sql.DB) *server {
	return &server{
		logger:   logger,
		cfg:      cfg,
		psqlConn: psqlConn,
	}
}

func (s *server) StartgRPC() {
	s.logger.Info("Starting gRPC server on port " + s.cfg.GrpcServer.Port)

	lis, err := net.Listen("tcp", s.cfg.GrpcServer.Port)
	if err != nil {
		s.logger.Error("Error while starting gRPC server: " + err.Error())
	}

	defer lis.Close()

	psqlRepo := repository.NewPostgresRepository(s.psqlConn, s.logger)
	cveUC := usecase.NewCveUC(psqlRepo)
	service := grpcservice.NewgRPCService(s.logger, s.cfg, cveUC)
	interceptors := interceptors.NewInterceptor(s.logger, s.cfg)

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.GrpcServer.MaxConnectionIdle * time.Minute,
			Timeout:           s.cfg.GrpcServer.Timeout * time.Second,
			MaxConnectionAge:  s.cfg.GrpcServer.MaxConnectionAge * time.Minute,
			Time:              s.cfg.GrpcServer.Timeout * time.Minute,
		}),
		grpc.ChainUnaryInterceptor(interceptors.Logger),
	}
	server := grpc.NewServer(opts...)
	cvepb.RegisterCveServiceServer(server, service)

	go func() {
		if err := server.Serve(lis); err != nil {
			s.logger.Error("Error while serving: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}
