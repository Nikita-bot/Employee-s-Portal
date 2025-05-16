package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	UserTaskRepository interface{}
	userTaskRepo       struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewUserTaskRepo(db *sqlx.DB, l *zap.Logger) UserTaskRepository {
	return &userTaskRepo{
		db: db,
		l:  l,
	}
}
