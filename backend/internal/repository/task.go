package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	TaskRepository interface {
		GetAll() ([]entity.Task, error)
	}

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

func (tr taskRepo) GetAll() ([]entity.Task, error) {
	tr.l.Info("IN TASK LIST REPO :: GET ALL TASKS")

	var tl []entity.Task

	query := `
		SELECT * from tasks
	`

	err := tr.db.Select(&tl, query)
	if err != nil {
		tr.l.Error(err.Error())
		return nil, err
	}

	return tl, nil
}
