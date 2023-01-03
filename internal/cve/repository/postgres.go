package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/diyliv/cve/internal/models"
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
	_, err := p.psql.Exec("INSERT INTO users (user_login, user_hashed_password) VALUES($1, $2)", creds.Login, creds.Password)
	if err != nil {
		p.logger.Error("Error while creating new user: " + err.Error())
		return err
	}

	return nil
}

func (p *postgresRepository) Login(ctx context.Context, creds models.Creds) (string, error) {
	var hashedPass string

	err := p.psql.QueryRow("SELECT user_hashed_password FROM users WHERE user_login = $1", creds.Login).Scan(hashedPass)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("this user doesnt exist")
		}
		p.logger.Error("Error while getting info about user: " + err.Error())
		return "", err
	}

	return hashedPass, nil
}
