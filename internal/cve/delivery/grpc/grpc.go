package grpc

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/diyliv/cve/config"
	"github.com/diyliv/cve/internal/cve"
	"github.com/diyliv/cve/internal/models"
	"github.com/diyliv/cve/pkg/errs"
	"github.com/diyliv/cve/pkg/hash"
	"github.com/diyliv/cve/pkg/parser"
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

	hashedPassword, err := hash.HashPass([]byte(data.Password))
	if err != nil {
		gs.logger.Error("Error while hashing password: " + err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := gs.cveUC.Register(ctx, models.Creds{Login: data.Login, Password: hashedPassword}); err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "This user already exists. Try another login.")
		}
		gs.logger.Error("Error while calling cveUC Register: " + err.Error())
	}

	return &cvepb.RegisterResp{Status: "created"}, status.Error(codes.OK, "")
}

func (gs *grpcservice) Login(ctx context.Context, req *cvepb.LoginReq) (*cvepb.LoginResp, error) {
	data := req.GetUserCreds()

	hashedPass, err := gs.cveUC.Login(ctx, data.Login)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		gs.logger.Error("Error while calling cveUC Login: " + err.Error())
	}

	if !hash.ComparePass(hashedPass, []byte(data.Password)) {
		return nil, status.Error(codes.Unauthenticated, "incorrect login/password")
	}

	return &cvepb.LoginResp{Status: "successful login"}, status.Error(codes.OK, "")
}

func (gs *grpcservice) FindVulnerabilities(ctx context.Context, req *cvepb.FindVulnerabilitiesReq) (*cvepb.FindVulnerabilitiesResp, error) {
	searchReq := req.GetSearch()

	parseResult := parser.ParseCveMitre(searchReq)
	if len(parseResult) == 0 {
		return nil, status.Error(codes.NotFound, errs.ErrNotFound.Error())
	}
	var resp []*cvepb.Cve

	for _, val := range parseResult {
		resp = append(resp, &cvepb.Cve{
			Name:        val.Name,
			Link:        val.Link,
			Description: val.Description,
		})
	}

	return &cvepb.FindVulnerabilitiesResp{Result: resp}, status.Error(codes.OK, "")
}
