package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	DepartmentRepository interface {
		GetUserOnDepartmentByTaskID(task_id int) ([]entity.User, error)
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

func (dr departmentRepo) GetUserOnDepartmentByTaskID(task_id int) ([]entity.User, error) {
	dr.l.Debug("IR DEPARTMENT REPO :: GET USERS ON DEPARTMENT")

	var u []entity.User

	query := `
		SELECT u.id, u.name, u.surname, u.patronymic, u.position 
		FROM users u
		JOIN task_department td ON u.department_id = td.department_id
		JOIN user_task ut ON ut.task_id = td.task_id
		WHERE ut.id = $1
	`

	err := dr.db.Select(&u, query, task_id)
	if err != nil {
		dr.l.Error(err.Error())
		return nil, err
	}

	return u, err

}
