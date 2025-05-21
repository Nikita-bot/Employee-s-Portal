package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type Handler struct {
	e                 *echo.Echo
	UserHandler       UserHandler
	UserTaskHandler   UserTaskHandler
	TaskHandler       TaskHandler
	RoleHandler       RoleHandler
	DepartmentHandler DepartmentHandler
	CommentHandler    CommentHandler
}

func NewHandler(l *zap.Logger, s service.Service, e *echo.Echo) *Handler {
	return &Handler{
		e:                 e,
		UserHandler:       NewUserHandler(s.UserService, l, e),
		UserTaskHandler:   NewUserTaskHandler(s.UserTaskService, l, e),
		TaskHandler:       NewTaskHandler(s.TaskService, l, e),
		RoleHandler:       NewRoleHandler(s.RoleService, l, e),
		DepartmentHandler: NewDepartmentHandler(s.DepartmentService, l, e),
		CommentHandler:    NewCommentHandler(s.CommentService, l, e),
	}
}

func (h *Handler) Handle(port string) error {

	h.UserHandler.Handle()
	h.UserTaskHandler.Handle()
	h.TaskHandler.Handle()
	h.RoleHandler.Handle()
	h.DepartmentHandler.Handle()
	h.CommentHandler.Handle()

	err := h.e.Start(":" + port)
	if err != nil {
		return err
	}
	return nil
}
