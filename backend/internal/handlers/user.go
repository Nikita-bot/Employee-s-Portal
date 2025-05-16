package handlers

import (
	"portal/internal/service"

	"go.uber.org/zap"
)

type (
	UserHandler interface{}
	userHandler struct {
		s service.UserService
		l *zap.Logger
	}
)

func NewUserHandler(s service.UserService, l *zap.Logger) UserHandler {
	return &userHandler{
		s: s,
		l: l,
	}
}
