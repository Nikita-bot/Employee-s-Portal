package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	TaskRepository interface{}

	taskRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewTaskRepo(db *sqlx.DB, l *zap.Logger) TaskRepository {
	return &taskRepo{
		db: db,
		l:  l,
	}
}
