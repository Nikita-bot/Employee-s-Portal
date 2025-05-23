package repository

import (
	"errors"
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	UserTaskRepository interface {
		TaskForUser(userId int) ([]entity.UserTask, error)
		TaskByUser(userId int) ([]entity.UserTask, error)
		CreateTask(uc entity.UserTaskCreate) error
		GetTaskByID(id int) (entity.UserTask, error)
		DeleteTask(id int) error
		UpdateTask(id int, date string) error
	}
	userTaskRepo struct {
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

func (u userTaskRepo) TaskForUser(userId int) ([]entity.UserTask, error) {
	u.l.Debug("IN USER TASK REPO :: GET TASK CREATED BY USER")

	var ut []entity.UserTask

	query := `
		SELECT * FROM user_task WHERE executor=$1 AND status!=3 AND status!=1
	`

	err := u.db.Select(&ut, query, userId)
	if err != nil {
		u.l.Error(err.Error())
		return nil, err
	}

	return ut, nil
}

func (u userTaskRepo) TaskByUser(userId int) ([]entity.UserTask, error) {
	u.l.Debug("IN USER TASK REPO :: GET TASK CREATED FOR USER")

	var ut []entity.UserTask

	query := `
		SELECT * FROM user_task WHERE initiator=$1 AND status!=3
	`

	err := u.db.Select(&ut, query, userId)
	if err != nil {
		u.l.Error(err.Error())
		return nil, err
	}

	return ut, nil
}

func (u userTaskRepo) CreateTask(uc entity.UserTaskCreate) error {
	u.l.Debug("IN USER TASK REPO :: CREATE TASK")

	_, err := u.db.NamedExec(`INSERT INTO user_task (task_id,executor,initiator, description,status,create_date,execute_date)
	VALUES (:task_id, :executor, :initiator, :description, :status, :create_date, :execute_date)`, uc)
	if err != nil {
		u.l.Error(err.Error())
		return err
	}

	return nil
}

func (u userTaskRepo) GetTaskByID(id int) (entity.UserTask, error) {
	u.l.Debug("IN USER TASK REPO :: GET TASK BY ID")

	query := `
		SELECT * FROM user_task WHERE id=$1 AND status!=3
	`

	var ut entity.UserTask

	err := u.db.Get(&ut, query, id)
	if err != nil {
		u.l.Error(err.Error())
		return entity.UserTask{}, err
	}

	return ut, nil
}

func (u userTaskRepo) DeleteTask(id int) error {
	u.l.Debug("IN USER TASK REPO :: DELETE TASK")

	query := `
		UPDATE user_task SET status=3 WHERE id=$1
	`

	_, err := u.db.Exec(query, id)
	if err != nil {
		u.l.Error(err.Error())
		return errors.New("Can't delete Task")
	}

	return nil
}

func (u userTaskRepo) UpdateTask(id int, date string) error {
	u.l.Debug("IN USER TASK REPO :: EXEC TASK")

	query := `
		UPDATE user_task SET status=1, execute_date = $2 WHERE id=$1
	`

	_, err := u.db.Exec(query, id, date)
	if err != nil {
		u.l.Error(err.Error())
		return errors.New("CAN`T EXEC TASK")
	}

	return nil
}
