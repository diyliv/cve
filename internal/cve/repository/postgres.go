package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type postgresRepository struct {
	psql   *sql.DB
	logger *zap.Logger
}

func NewPostgresRepository(psql *sql.DB, logger *zap.Logger) *postgresRepository {
	return &postgresRepository{
		psql:   psql,
		logger: logger,
	}
}
