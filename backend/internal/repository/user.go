package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	UserRepository interface {
		Login(login string) (string, error)
		GetUserByLogin(login string) (entity.User, error)
		GetUserByID(id int) (entity.User, error)
	}

	userRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewUserRepo(db *sqlx.DB, l *zap.Logger) UserRepository {
	return &userRepo{
		db: db,
		l:  l,
	}
}

func (u userRepo) Login(login string) (string, error) {
	u.l.Debug("IN USER REPO :: AUTH WITH LOGIN " + login)

	var password = ""

	query := `
		SELECT password FROM users WHERE login=$1
	`
	err := u.db.Get(&password, query, login)
	if err != nil {
		u.l.Debug(err.Error())
		return "", err
	}

	return password, nil
}

func (u userRepo) GetUserByLogin(login string) (entity.User, error) {
	u.l.Debug("IN USER REPO :: GET USER BY LOGIN " + login)

	var user entity.User

	query := `
		SELECT id,name,surname,patronymic,position
		,email,phone,tg_link,tg_id,pasport,snyls,
		adress,employment_date,dismissal_date
		FROM users WHERE login=$1
	`

	err := u.db.Get(&user, query, login)
	if err != nil {
		u.l.Error(err.Error())
		return entity.User{}, err
	}

	department := u.GetDepartmentByUserID(user.ID)
	user.Department = department
	boss := u.GetUserBoss(user.ID)
	user.Boss = boss

	return user, nil
}

func (u userRepo) GetDepartmentByUserID(id int) entity.Department {

	var department entity.Department

	query := `
		SELECT d.id, d.name FROM departments d, users u WHERE u.department_id = d.id AND u.id=$1
	`
	err := u.db.Get(&department, query, id)
	if err != nil {
		return entity.Department{}
	}

	return department

}

func (u userRepo) GetUserByID(id int) (entity.User, error) {
	u.l.Debug("IN USER REPO :: GET USER BY ID " + string(id))

	var user entity.User

	query := `
		SELECT id,name,surname,patronymic,position
		,email,phone,tg_link,tg_id,pasport,snyls,
		adress,employment_date,dismissal_date
		FROM users WHERE id=$1
	`

	err := u.db.Get(&user, query, id)
	if err != nil {
		u.l.Error(err.Error())
		return entity.User{}, err
	}

	department := u.GetDepartmentByUserID(user.ID)
	user.Department = department
	boss := u.GetUserBoss(user.ID)
	user.Boss = boss
	return user, nil
}

func (r userRepo) GetUserBoss(userID int) entity.UserMainData {

	var boss entity.UserMainData

	query := `
		SELECT u.id, u.name, u.surname, u.patronymic 
		FROM users u 
		JOIN users rab ON u.id = rab.boss 
		WHERE rab.id = $1
	`
	err := r.db.Get(&boss, query, userID)
	if err != nil {
		return entity.UserMainData{}
	}

	return boss
}
