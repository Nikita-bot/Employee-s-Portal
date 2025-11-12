package repository

import (
	"errors"
	"portal/internal/entity"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	UserTaskRepository interface {
		TaskForUser(userId int) ([]entity.UserTask, error)
		TaskByUser(userId int) ([]entity.UserTask, error)
		GetTaskByID(id int) (entity.UserTask, error)
		CreateTask(uc entity.UserTaskCreate) error
		DeleteTask(id int) error
		UpdateTask(id, status int, date string) error
		UpdateExecutor(id int, executor int) error

		GetUserAndCountTasksByRoleID(roleId, branchId, initiator int, isSuppot bool) ([]entity.UserCountTask, error)
		// GetDepartmentByTaskID(task_id int) []entity.Department
		GetTaskRoles(taskID int) (int, error)
	}
	userTaskRepo struct {
		db *sqlx.DB
		l  *zap.Logger
		r  *redis.Client
	}
)

func NewUserTaskRepo(db *sqlx.DB, l *zap.Logger, r *redis.Client) UserTaskRepository {
	return &userTaskRepo{
		db: db,
		l:  l,
		r:  r,
	}
}

func (u userTaskRepo) TaskForUser(userId int) ([]entity.UserTask, error) {
	u.l.Debug("IN USER TASK REPO :: GET TASK CREATED BY USER")

	var ut []entity.UserTask

	query := `
        SELECT 
            ut.id,
            ut.task_id,
			tl.name as task_name,
            ut.description,
            ut.status,
			ut.priority,
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
            tasks ut
        JOIN 
            users exec ON ut.executor = exec.id
		JOIN 
			task_list tl on tl.id = ut.task_id
        JOIN 
            users init ON ut.initiator = init.id
        WHERE 
            ut.executor = $1 
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

func (u userTaskRepo) TaskByUser(userId int) ([]entity.UserTask, error) {
	u.l.Debug("IN USER TASK REPO :: GET TASK CREATED FOR USER")

	var ut []entity.UserTask

	query := `
        SELECT
		ut.id,
		ut.task_id,
		tl.name as task_name,
		ut.description,
		ut.status,
		ut.priority,
		ut.create_date,
		ut.execute_date,
		exec.id AS "executor.id",
		exec.name AS "executor.name",
		exec.surname AS "executor.surname",
		exec.patronymic AS "executor.patronymic",
					-- Данные инициатора
		init.id AS "initiator.id",
		init.name AS "initiator.name",
		init.surname AS "initiator.surname",
		init.patronymic AS "initiator.patronymic"
		FROM tasks ut
		JOIN users exec ON ut.executor = exec.id
		JOIN users init ON ut.initiator = init.id
		join task_list tl on tl.id = ut.task_id
		WHERE ut.initiator = $1 AND ut.status != 3
		ORDER BY ut.create_date DESC
    `
	err := u.db.Select(&ut, query, userId)
	if err != nil {
		u.l.Error(err.Error())
		return nil, err
	}

	return ut, nil
}

// func (u userTaskRepo) GetDepartmentByTaskID(task_id int) []entity.Department {
// 	query := `
//         SELECT d.id, d.name
//         FROM departments d
//         JOIN task_department td ON d.id = td.department_id
//         WHERE td.task_id = $1
//     `

// 	var departments []entity.Department
// 	err := u.db.Select(&departments, query, task_id)
// 	if err != nil {
// 		u.l.Error(err.Error())
// 	}

// 	return departments
// }

func (u userTaskRepo) GetTaskRoles(taskID int) (int, error) {
	query := `
		SELECT role_id 
		FROM task_roles 
		WHERE task_id = $1
	`

	var roleID int
	err := u.db.Get(&roleID, query, taskID)
	if err != nil {
		u.l.Error("Failed to get task roles", zap.Error(err))
		return 0, err
	}

	u.l.Debug("Task roles", zap.Any("taskID", taskID), zap.Any("roleID", roleID))

	return roleID, nil
}

func (u userTaskRepo) GetUserAndCountTasksByRoleID(roleId, branchId, initiator int, isSupport bool) ([]entity.UserCountTask, error) {

	var uc []entity.UserCountTask

	type QueryParams struct {
		RoleId      int `db:"role_id"`
		BranchID    int `db:"branch_id"`
		InitiatorID int `db:"initiator_id"`
	}

	params := QueryParams{
		RoleId:      roleId,
		BranchID:    branchId,
		InitiatorID: initiator,
	}

	query := `
		SELECT u.id as user_id, COUNT(ut.id) as task_count, e.depart_id as user_department
		FROM users u
		LEFT JOIN tasks ut ON ut.executor = u.id and ut.status != 3
		left JOIN user_role ur ON ur.user_id = u.id
		JOIN employee e ON u.id = e.user_id
		JOIN branches b ON b.id = e.branch_id
		WHERE ur.role_id=:role_id and u.id != :initiator_id
    `

	if isSupport {
		query += " AND e.branch_id = :branch_id"
	}

	query += " GROUP BY u.id, e.depart_id"

	stmt, err := u.db.PrepareNamed(query)
	if err != nil {
		u.l.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Select(&uc, params)
	if err != nil {
		u.l.Error(err.Error())
		return nil, err
	}

	u.l.Debug("IN REPO USERS IN DEP", zap.Any("Users:", uc))

	return uc, nil
}

func (u userTaskRepo) CreateTask(uc entity.UserTaskCreate) error {
	u.l.Debug("IN USER TASK REPO :: CREATE TASK")

	var tj entity.TaskJournal
	var id int

	row := u.db.QueryRow(`
    INSERT INTO tasks 
    (task_id, executor, initiator, description, status, priority, create_date)
    VALUES ($1, $2, $3, $4, $5, $6, $7) 
    RETURNING id`,
		uc.Task, uc.Executor, uc.Initiator, uc.Description,
		uc.Status, uc.Priority, uc.CreateDate)

	err := row.Scan(&id)
	if err != nil {
		u.l.Error(err.Error())
		return err
	}

	loc := time.FixedZone("UTC+7", 7*60*60)
	tj = entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "Задача " + strconv.Itoa(int(id)) + " создана",
		Creation_date: time.Now().In(loc).Format("02.01.2006, 15:04:05"),
	}
	u.l.Debug("IN USER TASK REPO :: CREATE LOGGING IN JOURNAL")

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
			tl.name as task_name,
            ut.description,
            ut.status,
			ut.priority,
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
            tasks ut
        JOIN 
            users exec ON ut.executor = exec.id
		JOIN 
			task_list tl on tl.id = ut.task_id
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
		UPDATE tasks SET status=3 WHERE id=$1
	`

	_, err := u.db.Exec(query, id)
	if err != nil {
		u.l.Error(err.Error())
		return errors.New("can't delete Task")
	}

	loc := time.FixedZone("UTC+7", 7*60*60)
	tj := entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "Задача " + strconv.Itoa(int(id)) + " удалена",
		Creation_date: time.Now().In(loc).Format("02.01.2006, 15:04:05"),
	}

	_, err = u.db.NamedExec(`INSERT INTO task_journal (user_task_id, action, creation_date) VALUES (:user_task_id,:action,:creation_date)`, tj)
	if err != nil {
		u.l.Error(err.Error())
	}

	return nil
}

func (u userTaskRepo) UpdateTask(id, status int, date string) error {
	u.l.Debug("IN USER TASK REPO :: EXEC TASK")

	query := `
		UPDATE tasks SET status=$3, execute_date=$2 WHERE id=$1
	`

	_, err := u.db.Exec(query, id, date, status)
	if err != nil {
		u.l.Error(err.Error())
		return errors.New("CAN`T EXEC TASK")
	}

	var tj entity.TaskJournal
	loc := time.FixedZone("UTC+7", 7*60*60)

	if status == 2 {
		tj = entity.TaskJournal{
			UserTaskID:    int(id),
			Action:        "Задача " + strconv.Itoa(int(id)) + " выполнена",
			Creation_date: time.Now().In(loc).Format("02.01.2006, 15:04:05"),
		}
	} else if status == 1 {
		tj = entity.TaskJournal{
			UserTaskID:    int(id),
			Action:        "Задача " + strconv.Itoa(int(id)) + " взята в рабоу",
			Creation_date: time.Now().In(loc).Format("02.01.2006, 15:04:05"),
		}
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
		UPDATE tasks SET executor=$1 WHERE id=$2
	`

	_, err := u.db.Exec(query, executor, id)
	if err != nil {
		u.l.Error(err.Error())
		return errors.New("CAN`T CHANGE EXECUTOR")
	}

	loc := time.FixedZone("UTC+7", 7*60*60)
	tj := entity.TaskJournal{
		UserTaskID:    int(id),
		Action:        "У задачи " + strconv.Itoa(int(id)) + " изменен исполнитель",
		Creation_date: time.Now().In(loc).Format("02.01.2006, 15:04:05"),
	}

	_, err = u.db.NamedExec(`INSERT INTO task_journal (user_task_id, action, creation_date) VALUES (:user_task_id,:action,:creation_date)`, tj)
	if err != nil {
		u.l.Error(err.Error())
	}

	return nil
}
