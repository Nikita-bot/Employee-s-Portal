package handlers

import (
	"portal/internal/service"

	"go.uber.org/zap"
)

type (
	CommentHandler interface{}
	commentHandler struct {
		s service.CommentService
		l *zap.Logger
	}
)

func NewCommentHandler(s service.CommentService, l *zap.Logger) CommentHandler {
	return &commentHandler{
		s: s,
		l: l,
	}
}
