package domain

import (
	"context"
)

type EasyUser struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type EasyUserRepository interface {
	GetByUsernameAndPassword(context.Context, string, string) (*EasyUser, error)
}