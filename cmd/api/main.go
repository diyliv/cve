package main

import (
	"github.com/diyliv/cve/config"
	"github.com/diyliv/cve/internal/server"
	"github.com/diyliv/cve/pkg/logger"
	"github.com/diyliv/cve/pkg/storage/postgres"
)

func main() {
	cfg := config.ReadConfig()
	logger := logger.InitLogger()
	psqlConn := postgres.ConnectToPostgres(cfg)
	server := server.NewServer(logger, cfg, psqlConn)
	server.StartgRPC()
}
