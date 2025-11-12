package service

import (
	"fmt"
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	DepartmentService interface {
		GetUsersByTaskId(task_id int) ([]entity.User, error)
		GetAllDepartment() ([]entity.Department, error)
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

func (s departmentService) GetAllDepartment() ([]entity.Department, error) {
	s.l.Debug("IN SERVICE :: GET ALL Department")

	var n []entity.Department

	n, err := s.r.GetAll()
	if err != nil {
		e := fmt.Errorf("failed to get department: %w", err)
		return nil, e
	}

	return n, nil
}

func (ds departmentService) GetUsersByTaskId(task_id int) ([]entity.User, error) {
	ds.l.Debug("IN DEPARTMENT SERVISE :: GET USERS")

	u, err := ds.r.GetUserOnDepartmentByTaskID(task_id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
