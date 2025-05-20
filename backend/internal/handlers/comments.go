package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	CommentHandler interface {
		Handle()
	}
	commentHandler struct {
		s service.CommentService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewCommentHandler(s service.CommentService, l *zap.Logger, e *echo.Echo) CommentHandler {
	return &commentHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (ch commentHandler) Handle() {
	ch.e.GET("/comments", ch.getAllComments)
}

func (ch commentHandler) getAllComments(c echo.Context) error {
	return c.String(200, "Получи все коменты для задачи")
}
