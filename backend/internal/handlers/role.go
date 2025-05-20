package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	RoleHandler interface {
		Handle()
	}
	roleHandler struct {
		s service.RoleService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewRoleHandler(s service.RoleService, l *zap.Logger, e *echo.Echo) RoleHandler {
	return &roleHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (rh roleHandler) Handle() {
	rh.e.GET("/roles", rh.getAllRole)
}

func (rh roleHandler) getAllRole(c echo.Context) error {
	return c.String(200, "Получи все роли")
}
