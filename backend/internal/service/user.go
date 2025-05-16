package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	UserService interface{}
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
