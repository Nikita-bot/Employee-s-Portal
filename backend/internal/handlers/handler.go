package handlers

import (
	"net/http"
	"portal/internal/service"
	"portal/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Handler struct {
	conf              config.Config
	UserHandler       UserHandler
	UserTaskHandler   UserTaskHandler
	TaskHandler       TaskHandler
	RoleHandler       RoleHandler
	DepartmentHandler DepartmentHandler
	CommentHandler    CommentHandler
}

func NewHandler(l *zap.Logger, c config.Config, s service.Service) *Handler {
	return &Handler{
		conf:              c,
		UserHandler:       NewUserHandler(s.UserService, l),
		UserTaskHandler:   NewTaskHandler(s.UserTaskService, l),
		TaskHandler:       NewTaskHandler(s.TaskService, l),
		RoleHandler:       NewRoleHandler(s.RoleService, l),
		DepartmentHandler: NewDepartmentHandler(s.DepartmentService, l),
		CommentHandler:    NewCommentHandler(s.CommentService, l),
	}
}

func (h Handler) InitServer() error {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXRealIP,
			echo.HeaderXForwardedFor,
			echo.HeaderXRequestedWith,
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// e.GET("/api/v1/")

	err := e.Start(":" + h.conf.Port)
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}
