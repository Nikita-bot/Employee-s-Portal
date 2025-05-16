package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	CommetRepository interface{}
	commentRepo      struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewCommentRepo(db *sqlx.DB, l *zap.Logger) CommetRepository {
	return &commentRepo{
		db: db,
		l:  l,
	}
}
