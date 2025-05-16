package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	RoleRepository interface{}
	roleRepo       struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewRoleRepo(db *sqlx.DB, l *zap.Logger) RoleRepository {
	return &roleRepo{
		db: db,
		l:  l,
	}
}
