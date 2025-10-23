package service

import (
	"errors"
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	UserService interface {
		Login(login, pass string) (int, error)
		GetUserByID(id int) (entity.UserMainData, error)
		ChangePassword(id int, pass string) error
	}
	userService struct {
		r repository.UserRepository
		l *zap.Logger
	}
)

func NewUserService(r repository.UserRepository, l *zap.Logger) UserService {
	return &userService{
		r: r,
		l: l,
	}
}

func (u userService) Login(login, pass string) (int, error) {
	u.l.Debug("IN SERVICE :: AUTH USER " + login)

	dataBasePass, err := u.r.Login(login)
	if err != nil {
		return 0, err
	}
	if dataBasePass == "" {
		dataBasePass = pass
		err = u.r.CreatePassword(login, pass)
		if err != nil {
			return 0, err
		}

	}

	if pass != dataBasePass {
		return 0, errors.New("WRONG LOGIN OR PASSWORD")
	}

	id, err := u.r.GetUserByLogin(login)
	if err != nil {
		return 0, errors.New("ERROR WHEN TRY GET USER BY LOGIN")
	}

	return id, nil
}

func (u userService) GetUserByID(id int) (entity.UserMainData, error) {
	u.l.Debug("IN SERVICE :: GET USER BY ID")

	var user entity.UserMainData

	user, err := u.r.GetUserByID(id)
	if err != nil {
		return entity.UserMainData{}, errors.New("ERROR WHEN GET USER BY ID")
	}

	return user, nil

}

func (u userService) ChangePassword(id int, pass string) error {
	u.l.Debug("IN USER REPO :: CHANGE PASSWORD")

	err := u.r.ChangeUserPassword(id, pass)
	if err != nil {
		return err
	}

	return nil
}
