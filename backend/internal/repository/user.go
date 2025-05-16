package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	UserRepository interface {
	}

	userRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewUserRepo(db *sqlx.DB, l *zap.Logger) UserTaskRepository {
	return &userRepo{
		db: db,
		l:  l,
	}
}
