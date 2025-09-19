package service

import (
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	TaskService interface {
		GetAllTask() ([]entity.Task, error)
		GetITTask() ([]entity.Task, error)
		GetTaskByID(id int) (entity.Task, error)
	}
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

func (ts taskService) GetAllTask() ([]entity.Task, error) {
	ts.l.Debug("IN TASK LIST SERVICE :: GET ALL TASK")

	tl, err := ts.r.GetAll()
	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (ts taskService) GetITTask() ([]entity.Task, error) {
	ts.l.Debug("IN TASK LIST SERVICE :: GET IT TASK")

	tl, err := ts.r.GetIT()
	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (ts taskService) GetTaskByID(id int) (entity.Task, error) {
	ts.l.Debug("IN TASK LIST SERVICE :: GET TASK BY ID")

	t, err := ts.r.GetByID(id)
	if err != nil {
		return entity.Task{}, err
	}

	return t, nil
}
