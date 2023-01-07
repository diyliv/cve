package grpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/diyliv/cve/config"
	"github.com/diyliv/cve/internal/cve"
	"github.com/diyliv/cve/internal/models"
	"github.com/diyliv/cve/pkg/hash"
	cvepb "github.com/diyliv/cve/proto/cve"
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

func (gs *grpcservice) Register(ctx context.Context, req *cvepb.RegisterReq) (*cvepb.RegisterResp, error) {
	data := req.GetUserCreds()
	var creds models.Creds

	creds.Login = data.Login
	hashedPassword, err := hash.HashPass([]byte(creds.Password))
	if err != nil {
		gs.logger.Error("Error while hashing password: " + err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	creds.Password = hashedPassword

	if err := gs.cveUC.Register(ctx, creds); err != nil {
		return nil, status.Error(codes.AlreadyExists, "This user already exists. Try another login.")
	}

	return &cvepb.RegisterResp{Status: "created"}, status.Error(codes.OK, "")
}

func (gs *grpcservice) Login(ctx context.Context, req *cvepb.LoginReq) (*cvepb.LoginResp, error) {
	return nil, nil
}

func (gs *grpcservice) FindVulnerabilities(ctx context.Context, req *cvepb.FindVulnerabilitiesReq) (*cvepb.FindVulnerabilitiesResp, error) {
	return nil, nil
}
