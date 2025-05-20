package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	UserTaskHandler interface {
		Handle()
	}
	userTaskHandler struct {
		s service.UserTaskService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewUserTaskHandler(s service.UserTaskService, l *zap.Logger, e *echo.Echo) UserTaskHandler {
	return &userTaskHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (uh userTaskHandler) Handle() {
	uh.e.GET("/userTask", uh.getAllUserTask)
}

func (uh userTaskHandler) getAllUserTask(c echo.Context) error {
	return c.String(200, "Получи все таски для Юсера")
}
