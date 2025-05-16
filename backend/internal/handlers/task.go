package handlers

import (
	"portal/internal/service"

	"go.uber.org/zap"
)

type (
	TaskHandler interface{}
	taskHandler struct {
		s service.TaskService
		l *zap.Logger
	}
)

func NewTaskHandler(s service.TaskService, l *zap.Logger) TaskHandler {
	return &taskHandler{
		s: s,
		l: l,
	}
}
