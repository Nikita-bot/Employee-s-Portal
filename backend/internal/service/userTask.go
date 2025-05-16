package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	UserTaskService interface{}
	userTaskService struct {
		r repository.UserTaskRepository
		l *zap.Logger
	}
)

func NewUserTaskService(r repository.UserTaskRepository, l *zap.Logger) UserTaskService {
	return &userTaskService{
		r: r,
		l: l,
	}
}
