package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	DepartmentHandler interface {
		Handle()
	}
	departmentHandler struct {
		s service.DepartmentService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewDepartmentHandler(s service.DepartmentService, l *zap.Logger, e *echo.Echo) DepartmentHandler {
	return &departmentHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (dh departmentHandler) Handle() {
	dh.e.GET("/depatments", dh.getAllDepartments)
}

func (dh departmentHandler) getAllDepartments(c echo.Context) error {
	return c.String(200, "Получи все отделения")
}
