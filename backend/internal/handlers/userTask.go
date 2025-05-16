package handlers

import (
	"portal/internal/service"

	"go.uber.org/zap"
)

type (
	UserTaskHandler interface{}
	userTaskHandler struct {
		s service.UserTaskService
		l *zap.Logger
	}
)

func NewUserTaskHandler(s service.UserTaskService, l *zap.Logger) UserTaskHandler {
	return &userTaskHandler{
		s: s,
		l: l,
	}
}
