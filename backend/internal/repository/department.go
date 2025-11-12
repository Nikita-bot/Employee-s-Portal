package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	DepartmentRepository interface {
		GetUserOnDepartmentByTaskID(task_id int) ([]entity.User, error)
		GetAll() ([]entity.Department, error)
	}
	departmentRepo struct {
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

func (r departmentRepo) GetAll() ([]entity.Department, error) {
	r.l.Info("IN NEWS REPO :: GET ALL Department")

	var n []entity.Department

	query := `
        SELECT 
            n.id, 
            n.name
        FROM departments n
    `
	err := r.db.Select(&n, query)
	if err != nil {
		r.l.Error(err.Error())
		return nil, err
	}
	return n, nil
}

func (dr departmentRepo) GetUserOnDepartmentByTaskID(task_id int) ([]entity.User, error) {
	dr.l.Debug("IN DEPARTMENT REPO :: GET USERS ON DEPARTMENT")

	var u []entity.User

	query := `
		SELECT u.id, u.name, u.surname, u.patronymic
		FROM users u
		JOIN employee e ON e.user_id = u.id
		JOIN task_department td ON e.depart_id = td.department_id
		JOIN tasks ut ON ut.task_id = td.task_id
		WHERE ut.id = $1
	`

	err := dr.db.Select(&u, query, task_id)
	if err != nil {
		dr.l.Error(err.Error())
		return nil, err
	}

	return u, err

}
