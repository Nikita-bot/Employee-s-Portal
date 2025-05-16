package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	DepartmentService interface{}
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
