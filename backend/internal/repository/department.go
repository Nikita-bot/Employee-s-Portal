package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	DepartmentRepository interface{}
	departmentRepo       struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewDepartmentRepo(db *sqlx.DB, l *zap.Logger) DepartmentRepository {
	return &departmentRepo{
		db: db,
		l:  l,
	}
}
