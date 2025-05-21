package handlers

import (
	"portal/internal/entity"
	"portal/internal/service"
	"strconv"

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
	ch.e.POST("/comments", ch.createComment)
	ch.e.GET("/api/v1/comments/:taskId", ch.getAllComments)

}

func (ch commentHandler) getAllComments(c echo.Context) error {
	ch.l.Debug("IN COMMENT HANDLER :: GET ALL COMMENTS")

	var response = make(map[string]interface{})

	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		return c.String(401, "ID MUST BE INT")
	}

	cl, err := ch.s.GetComments(taskId)
	if err != nil {
		return c.String(501, err.Error())
	}

	response["comments"] = cl

	return c.JSON(200, response)
}

func (ch commentHandler) createComment(c echo.Context) error {
	ch.l.Debug("IN COMMENT HANDLER :: CREATE NEW COMMENTS")

	var comment entity.CommentCreate

	err := c.Bind(&comment)
	if err != nil {
		return c.String(401, "BAD REQUEST")
	}

	err = ch.s.CreateComment(comment)
	if err != nil {
		return c.String(501, err.Error())
	}

	return c.String(200, "Ok")
}
