package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	UserHandler interface {
		Handle()
	}
	userHandler struct {
		s service.UserService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewUserHandler(s service.UserService, l *zap.Logger, e *echo.Echo) UserHandler {
	return &userHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (uh userHandler) Handle() {
	uh.e.GET("/users", uh.getAllUser)
}

func (uh userHandler) getAllUser(c echo.Context) error {
	return c.String(200, "Получи всех Юсеров")
}
