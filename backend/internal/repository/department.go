package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	DepartmentRepository interface {
		GetUserOnDepartment(dep_id int) ([]entity.User, error)
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

func (dr departmentRepo) GetUserOnDepartment(dep_id int) ([]entity.User, error) {
	dr.l.Debug("IR DEPARTMENT REPO :: GET USERS ON DEPARTMENT")

	var u []entity.User

	query := `
		SELECT id, name, surname, patronymic, position  FROM users WHERE department_id=$1
	`

	err := dr.db.Select(&u, query, dep_id)
	if err != nil {
		dr.l.Error(err.Error())
		return nil, err
	}

	return u, err

}
