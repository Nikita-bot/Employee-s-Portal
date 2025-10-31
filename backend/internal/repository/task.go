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
		GetAvailableUsers(task_id int) ([]entity.UserMainData, error)
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

func (tr taskRepo) GetAvailableUsers(task_id int) ([]entity.UserMainData, error) {
	tr.l.Info("IN TASK LIST REPO :: GET AVAILABLE USERS")

	var users []entity.UserMainData
	query := `
		SELECT 
			u.id,
			u.name,
			u.surname,
			u.patronymic,
			u.phone,
			u.email,
			u.tg_link,
			u.tg_id
		FROM users u
		join user_role ur on u.id = ur.user_id
		join task_roles tr on ur.role_id = tr.role_id
		WHERE tr.task_id = $1 
		`
	err := tr.db.Select(&users, query, task_id)
	if err != nil {
		tr.l.Error(err.Error())
		return nil, err
	}
	return users, nil
}

func (tr taskRepo) GetAll() ([]entity.Task, error) {
	tr.l.Info("IN TASK LIST REPO :: GET ALL TASKS")

	var tl []entity.Task

	query := `
		SELECT * from task_list WHERE type != 'support'
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
		SELECT * FROM task_list WHERE id=$1
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
		from task_list t
		WHERE t.type = 'support'
	`

	err := tr.db.Select(&tl, query)
	if err != nil {
		tr.l.Error(err.Error())
		return nil, err
	}

	return tl, nil
}
