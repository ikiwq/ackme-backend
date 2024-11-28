package repository

import (
	"context"
	"fmt"

	"github.com/ikiwq/ackme/easy-api/internal/domain"
	"github.com/jmoiron/sqlx"
)

type mysqlEasyUserRepository struct {
	db *sqlx.DB
}

func NewMySqlEasyUserRepository(db *sqlx.DB) domain.EasyUserRepository {
	return &mysqlEasyUserRepository{db: db}
}

func (p *mysqlEasyUserRepository) GetByUsernameAndPassword(ctx context.Context, username string, password string) (*domain.EasyUser, error) {
	// This is made purposely unsafe and this is meant to BE a vulnerability. Do NOT make this more secure!
	query := "SELECT * FROM user WHERE username = '" + username + "' AND password = '" + password + "'";

	user := domain.EasyUser{}
	if err := p.db.GetContext(ctx, &user, query); err != nil {
		return nil, err
	}

	fmt.Println(user)

	return &user, nil
}