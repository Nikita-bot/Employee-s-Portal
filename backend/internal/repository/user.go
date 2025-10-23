package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"portal/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	UserRepository interface {
		Login(login string) (string, error)
		CreatePassword(login, pass string) error
		ChangeUserPassword(userID int, newPassword string) error
		GetUserByLogin(login string) (int, error)
		GetUserByID(id int) (entity.UserMainData, error)
	}

	userRepo struct {
		db *sqlx.DB
		l  *zap.Logger
		r  *redis.Client
	}
)

func NewUserRepo(db *sqlx.DB, l *zap.Logger, r *redis.Client) UserRepository {
	return &userRepo{
		db: db,
		l:  l,
		r:  r,
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

func (u userRepo) CreatePassword(login, pass string) error {
	u.l.Debug("IN USER REPO :: AUTH WITH LOGIN ")

	query := `
        UPDATE users 
        SET password = $1
        WHERE login = $2
    `

	_, err := u.db.Exec(query, pass, login)
	if err != nil {
		u.l.Error("failed to create password: " + err.Error())
		return err
	}

	return nil
}

func (u userRepo) ChangeUserPassword(userID int, newPassword string) error {

	u.l.Debug("IN USER REPO :: CHANGE PASSWORD")

	u.l.Debug("Fields:", zap.Int("ID", userID), zap.String("Password", newPassword))

	query := `
        UPDATE users 
        SET password = $1
        WHERE id = $2 
    `

	_, err := u.db.Exec(query, newPassword, userID)
	if err != nil {
		u.l.Error("failed to create password: " + err.Error())
		return err
	}

	return nil
}

func (u userRepo) GetUserByLogin(login string) (int, error) {
	u.l.Debug("IN USER REPO :: GET USER BY LOGIN " + login)

	var id int

	query := `
		SELECT 
			u.id
		FROM 
			users u
		WHERE u.login=$1 
		`

	err := u.db.Get(&id, query, login)
	if err != nil {
		u.l.Error(err.Error())
		return 0, err
	}

	return id, nil
}

func (u userRepo) GetUserRoles(id int) []entity.Role {
	u.l.Debug("IN USER REPO :: GET USER ROLES ")
	var r []entity.Role

	query := `
		SELECT r.id, r.name FROM roles r
		JOIN user_role ur ON ur.role_id=r.id
		WHERE ur.user_id=$1
	`

	err := u.db.Select(&r, query, id)
	if err != nil {
		return []entity.Role{}
	}

	return r
}

func (u userRepo) GetDepartmentByUserID(id int) entity.Department {

	var department entity.Department

	query := `
		SELECT d.id, d.name 
		FROM departments d
		JOIN employee e ON e.depart_id = d.id
		JOIN users u ON u.id = e.user_id
		WHERE u.id = $1
	`
	err := u.db.Get(&department, query, id)
	if err != nil {
		return entity.Department{}
	}

	return department

}

func (u userRepo) GetBranchByUserID(id int) entity.Branch {

	var department entity.Branch

	query := `
		SELECT d.id, d.name 
		FROM branches d
		JOIN employee e ON e.branch_id = d.id
		JOIN users u ON u.id = e.user_id
		WHERE u.id = $1
	`
	err := u.db.Get(&department, query, id)
	if err != nil {
		return entity.Branch{}
	}

	return department

}

func (u userRepo) GetUserByID(id int) (entity.UserMainData, error) {
	u.l.Debug("IN USER REPO :: GET USER BY ID ")

	var user entity.UserMainData

	cacheKey := fmt.Sprintf("user:%d", id)

	if data, err := u.r.Get(context.Background(), cacheKey).Bytes(); err == nil {
		if err := json.Unmarshal(data, &user); err == nil {
			u.l.Debug("Пользователь найден в Кеше")
			return user, nil
		}
	}

	u.l.Debug("Пользователь не найден в Кеше. Поиск в БД")

	query := `
		SELECT 
			u.id,
			u.name,
			u.surname,
			u.patronymic,
			u.phone,
			u.email,
			u.tg_link,
			u.tg_id,
			e.id AS "employee.id",
			e.tab_num AS "employee.tab_num",
			e.user_id AS "employee.user_id",
			e.zanyatost AS "employee.zanyatost",
			e.start_date AS "employee.start_date",
			e.end_date AS "employee.end_date",
			e.position AS "employee.position",
			e.depart_id AS "employee.depart_id",
			e.branch_id AS "employee.branch_id"
		FROM 
			users u
		LEFT JOIN 
			employee e ON u.id = e.user_id 
			AND (e.zanyatost = 'Основное место работы' OR e.zanyatost = 'Внешнее совместительство')
		WHERE 
			u.id = $1
		LIMIT 1
	`

	err := u.db.Get(&user, query, id)
	if err != nil {
		u.l.Error(err.Error())
		return entity.UserMainData{}, err
	}

	user.Employee.OtherPosition = []entity.Position{}
	department := u.GetDepartmentByUserID(user.ID)
	user.Employee.Department = department.Name
	branch := u.GetBranchByUserID(user.ID)
	user.Employee.Branch = branch.Name
	positions := u.GetAllUserPosition(user.ID)
	user.Employee.OtherPosition = append(user.Employee.OtherPosition, positions...)

	if data, err := json.Marshal(user); err == nil {
		u.r.Set(context.Background(), cacheKey, data, time.Hour)
	}

	return user, nil
}

// func (r userRepo) GetUserBoss(userID int) entity.UserMainData {

// 	var boss entity.UserMainData

// 	query := `
// 		SELECT u.name, u.surname, u.patronymic
// 		FROM users u
// 		WHERE u.id = $1
// 	`
// 	err := r.db.Get(&boss, query, userID)
// 	if err != nil {
// 		return entity.UserMainData{}
// 	}

// 	return boss
// }

func (r userRepo) GetAllUserPosition(userID int) []entity.Position {
	r.l.Info("IN USER REPO :: GET USER`S POSITION")

	var pos []entity.Position

	query := `
	SELECT e.position as "name", d.name as "department" 
	FROM employee e 
	JOIN departments d ON e.depart_id = d.id
	WHERE e.user_id = $1
	AND (e.zanyatost != 'Основное место работы' OR e.zanyatost != 'Внешнее совместительство') 
	`
	err := r.db.Select(&pos, query, userID)
	if err != nil {
		return nil
	}

	return pos
}
