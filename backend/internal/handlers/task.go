package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	TaskHandler interface {
		Handle()
	}
	taskHandler struct {
		s service.TaskService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewTaskHandler(s service.TaskService, l *zap.Logger, e *echo.Echo) TaskHandler {
	return &taskHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (th taskHandler) Handle() {
	th.e.GET("/tasks", th.getAllTask)
}

func (th taskHandler) getAllTask(c echo.Context) error {
	return c.String(200, "Получи все таски ")
}
