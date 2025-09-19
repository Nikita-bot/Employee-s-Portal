package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	TaskRepository interface {
		GetAll() ([]entity.Task, error)
		GetIT() ([]entity.Task, error)
		GetByID(id int) (entity.Task, error)
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

func (tr taskRepo) GetByID(id int) (entity.Task, error) {
	tr.l.Info("IN TASK LIST REPO :: GET ALL TASKS")

	var tl entity.Task

	query := `
		SELECT * FROM tasks WHERE id=$1
	`

	err := tr.db.Get(&tl, query, id)
	if err != nil {
		tr.l.Error(err.Error())
		return entity.Task{}, err
	}

	return tl, nil
}

func (tr taskRepo) GetIT() ([]entity.Task, error) {
	tr.l.Info("IN TASK LIST REPO :: GET IT TASKS")

	var tl []entity.Task
	//449 - отдел АСУ
	query := `
		SELECT t.id, t.name, t.type
		from tasks t
		join task_department td on td.task_id = t.id
		WHERE td.department_id = 449
	`

	err := tr.db.Select(&tl, query)
	if err != nil {
		tr.l.Error(err.Error())
		return nil, err
	}

	return tl, nil
}
