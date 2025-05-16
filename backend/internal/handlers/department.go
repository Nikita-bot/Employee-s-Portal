package handlers

import (
	"portal/internal/service"

	"go.uber.org/zap"
)

type (
	DepartmentHandler interface{}
	departmentHandler struct {
		s service.DepartmentService
		l *zap.Logger
	}
)

func NewDepartmentHandler(s service.DepartmentService, l *zap.Logger) DepartmentHandler {
	return &departmentHandler{
		s: s,
		l: l,
	}
}
