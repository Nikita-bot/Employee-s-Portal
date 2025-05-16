package handlers

import (
	"portal/internal/service"

	"go.uber.org/zap"
)

type (
	RoleHandler interface{}
	roleHandler struct {
		s service.RoleService
		l *zap.Logger
	}
)

func NewRoleHandler(s service.RoleService, l *zap.Logger) RoleHandler {
	return &roleHandler{
		s: s,
		l: l,
	}
}
