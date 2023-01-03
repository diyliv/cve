package server

import (
	"database/sql"

	"github.com/diyliv/cve/config"
	"go.uber.org/zap"
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

}
