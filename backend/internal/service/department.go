package service

import (
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	DepartmentService interface {
		GetUsersByTaskId(task_id int) ([]entity.User, error)
	}
	departmentService struct {
		r repository.DepartmentRepository
		l *zap.Logger
	}
)

func NewDepartmentService(r repository.DepartmentRepository, l *zap.Logger) DepartmentService {
	return &departmentService{
		r: r,
		l: l,
	}
}

func (ds departmentService) GetUsersByTaskId(task_id int) ([]entity.User, error) {
	ds.l.Debug("IN DEPARTMENT SERVISE :: GET USERS")

	u, err := ds.r.GetUserOnDepartmentByTaskID(task_id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
