package service

import (
	"errors"
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	UserService interface {
		Login(login, pass string) (entity.User, error)
		GetUserByID(id int) (entity.User, error)
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

func (u userService) Login(login, pass string) (entity.User, error) {
	u.l.Debug("IN SERVICE :: AUTH USER " + login)

	dataBasePass, err := u.r.Login(login)
	if err != nil {
		return entity.User{}, err
	}

	if pass != dataBasePass {
		return entity.User{}, errors.New("WRONG LOGIN OR PASSWORD")
	}

	user, err := u.r.GetUserByLogin(login)
	if err != nil {
		return entity.User{}, errors.New("ERROR WHEN TRY GET USER BY LOGIN")
	}
	return user, nil
}

func (u userService) GetUserByID(id int) (entity.User, error) {
	u.l.Debug("IN SERVICE :: GET USER BY ID" + string(id))

	var user entity.User

	user, err := u.r.GetUserByID(id)
	if err != nil {
		return entity.User{}, errors.New("ERROR WHEN GET USER BY ID")
	}

	return user, nil

}
