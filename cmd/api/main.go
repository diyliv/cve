package main

import (
	"github.com/diyliv/cve/config"
	"github.com/diyliv/cve/pkg/storage/postgres"
)

func main() {
	cfg := config.ReadConfig()
	postgres.ConnectToPostgres(cfg)
}
