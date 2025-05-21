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
	uh.e.GET("/api/v1/tasks/:userId", uh.getAllUserTask)
	uh.e.POST("/api/v1/tasks", uh.postUserTask)
	uh.e.GET("/api/v1/tasks/:id", uh.getUserTaskByID)
}

func (uh userTaskHandler) getAllUserTask(c echo.Context) error {
	return c.String(200, "Получи все таски для Юсера")
}

func (uh userTaskHandler) postUserTask(c echo.Context) error {
	return c.String(200, "Получи все таски ")
}

func (uh userTaskHandler) getUserTaskByID(c echo.Context) error {
	return c.String(200, "Получи все таски ")
}
