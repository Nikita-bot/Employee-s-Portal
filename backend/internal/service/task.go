package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	TaskService interface{}
	taskService struct {
		r repository.TaskRepository
		l *zap.Logger
	}
)

func NewTaskService(r repository.TaskRepository, l *zap.Logger) TaskService {
	return &taskService{
		r: r,
		l: l,
	}
}
