package repository

import (
	"errors"
	"portal/internal/entity"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	UserTaskRepository interface {
		TaskForUser(userId int) ([]entity.UserTask, error)
		TaskByUser(userId int) ([]entity.UserTask, error)
		GetUserAndCountTasksByTaskID(id int) ([]entity.UserCountTask, error)
		CreateTask(uc entity.UserTaskCreate) error
		GetTaskByID(id int) (entity.UserTask, error)
		DeleteTask(id int) error
		UpdateTask(id int, date string) error
		UpdateExecutor(id int, executor int) error
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
        SELECT 
            ut.id,
            ut.task_id,
            ut.description,
            ut.status,
            ut.create_date,
            ut.execute_date,
            -- Данные исполнителя
            exec.id AS "executor.id",
            exec.name AS "executor.name",
            exec.surname AS "executor.surname",
            exec.patronymic AS "executor.patronymic",
            -- Данные инициатора
            init.id AS "initiator.id",
            init.name AS "initiator.name",
            init.surname AS "initiator.surname",
            init.patronymic AS "initiator.patronymic"
        FROM 
            user_task ut
        JOIN 
            users exec ON ut.executor = exec.id
        JOIN 
            users init ON ut.initiator = init.id
        WHERE 
            ut.executor = $1 
            AND ut.status NOT IN (1, 3)
        ORDER BY 
            ut.create_date DESC
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
        SELECT 
            ut.id,
            ut.task_id,
            ut.description,
            ut.status,
            ut.create_date,
            ut.execute_date,
            -- Данные исполнителя
            exec.id AS "executor.id",
            exec.name AS "executor.name",
            exec.surname AS "executor.surname",
            exec.patronymic AS "executor.patronymic",
            -- Данные инициатора
            init.id AS "initiator.id",
            init.name AS "initiator.name",
            init.surname AS "initiator.surname",
            init.patronymic AS "initiator.patronymic"
        FROM 
            user_task ut
        JOIN 
            users exec ON ut.executor = exec.id
        JOIN 
            users init ON ut.initiator = init.id
        WHERE 
            ut.initiator = $1 
            AND ut.status != 3
        ORDER BY 
            ut.create_date DESC
    `

	err := u.db.Select(&ut, query, userId)
	if err != nil {
		u.l.Error(err.Error())
		return nil, err
	}

	return ut, nil
}

/*
Получение id пользователей и кол-во задач по task_id

SELECT u.id as user_id,COUNT(ut.id) as task_count
FROM users u
LEFT JOIN user_task ut ON u.id = ut.executor AND ut.status != 3
join tasks t  on t.id  = $1 and t.department_id = u.department_id
GROUP BY u.id
*/

func (u userTaskRepo) GetUserAndCountTasksByTaskID(id int) ([]entity.UserCountTask, error) {

	var uc []entity.UserCountTask

	query := `
		SELECT u.id as user_id,COUNT(ut.id) as task_count
		FROM users u
		LEFT JOIN user_task ut ON u.id = ut.executor AND ut.status != 3
		join tasks t  on t.id  = $1 and t.department_id = u.department_id
		GROUP BY u.id
	`

	err := u.db.Select(&uc, query, id)
	if err != nil {
		u.l.Error(err.Error())
		return []entity.UserCountTask{}, err
	}

	return uc, nil
}

func (u userTaskRepo) CreateTask(uc entity.UserTaskCreate) error {
	u.l.Debug("IN USER TASK REPO :: CREATE TASK")

	var tj entity.TaskJournal
	var id int

	row := u.db.QueryRow(`
    INSERT INTO user_task 
    (task_id, executor, initiator, description, status, create_date)
    VALUES ($1, $2, $3, $4, $5, $6) 
    RETURNING id`,
		uc.Task, uc.Executor, uc.Initiator, uc.Description,
		uc.Status, uc.CreateDate)

	err := row.Scan(&id)
	if err != nil {
		u.l.Error(err.Error())
		return err
	}

	tj = entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "Задача " + strconv.Itoa(int(id)) + " создана",
		Creation_date: time.Now().String(),
	}

	_, err = u.db.NamedExec(`INSERT INTO task_journal (user_task_id, action, creation_date) VALUES (:user_task_id,:action,:creation_date)`, tj)
	if err != nil {
		u.l.Error(err.Error())
	}
	return nil
}

func (u userTaskRepo) GetTaskByID(id int) (entity.UserTask, error) {
	u.l.Debug("IN USER TASK REPO :: GET TASK BY ID")

	query := `
        SELECT 
            ut.id,
            ut.task_id,
            ut.description,
            ut.status,
            ut.create_date,
            ut.execute_date,
            -- Данные исполнителя
            exec.id AS "executor.id",
            exec.name AS "executor.name",
            exec.surname AS "executor.surname",
            exec.patronymic AS "executor.patronymic",
            -- Данные инициатора
            init.id AS "initiator.id",
            init.name AS "initiator.name",
            init.surname AS "initiator.surname",
            init.patronymic AS "initiator.patronymic"
        FROM 
            user_task ut
        JOIN 
            users exec ON ut.executor = exec.id
        JOIN 
            users init ON ut.initiator = init.id
        WHERE 
            ut.id = $1 
            AND ut.status != 3
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

	tj := entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "Задача " + strconv.Itoa(int(id)) + " удалена",
		Creation_date: time.Now().String(),
	}

	_, err = u.db.NamedExec(`INSERT INTO task_journal (user_task_id, action, creation_date) VALUES (:user_task_id,:action,:creation_date)`, tj)
	if err != nil {
		u.l.Error(err.Error())
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

	tj := entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "Задача " + strconv.Itoa(int(id)) + " выполнена",
		Creation_date: time.Now().String(),
	}

	_, err = u.db.NamedExec(`INSERT INTO task_journal (user_task_id, action, creation_date) VALUES (:user_task_id,:action,:creation_date)`, tj)
	if err != nil {
		u.l.Error(err.Error())
	}

	return nil
}

func (u userTaskRepo) UpdateExecutor(id int, executor int) error {
	u.l.Debug("IN USER TASK REPO :: CHANGE EXECUTOR TASK")

	query := `
		UPDATE user_task SET executor=$1 WHERE id=$2
	`

	_, err := u.db.Exec(query, executor, id)
	if err != nil {
		u.l.Error(err.Error())
		return errors.New("CAN`T CHANGE EXECUTOR")
	}

	tj := entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "У задачи " + strconv.Itoa(int(id)) + " изменен исполнитель",
		Creation_date: time.Now().String(),
	}

	_, err = u.db.NamedExec(`INSERT INTO task_journal (user_task_id, action, creation_date) VALUES (:user_task_id,:action,:creation_date)`, tj)
	if err != nil {
		u.l.Error(err.Error())
	}

	return nil
}
