package repository

import (
	"context"
	"database/sql"

	"github.com/diyliv/cve/internal/models"
	"github.com/diyliv/cve/pkg/errs"
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

func (p *postgresRepository) Register(ctx context.Context, creds models.Creds) error {
	rows, err := p.psql.Exec("INSERT INTO users (user_login, user_hashed_password) VALUES($1, $2)", creds.Login, creds.Password)
	if rows == nil {
		return errs.ErrAlreadyExists
	}
	if err != nil {
		p.logger.Error("Error while creating new user: " + err.Error())
		return err
	}

	return nil
}

func (p *postgresRepository) Login(ctx context.Context, login string) (string, error) {
	var hashedPass string

	err := p.psql.QueryRow("SELECT user_hashed_password FROM users WHERE user_login = $1", login).Scan(&hashedPass)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errs.ErrNotFound
		}
		p.logger.Error("Error while getting info about user: " + err.Error())
		return "", err
	}

	return hashedPass, nil
}
