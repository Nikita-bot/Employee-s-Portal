package handlers

import (
	"portal/internal/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	JournalHandler    JournalHandler
	NewsHandler       NewsHandler
	PrinterHandler    PrinterHandler
	RoomHandler       RoomHandler
}

func NewHandler(l *zap.Logger, s service.Service, e *echo.Echo, secret string) *Handler {
	return &Handler{
		e:                 e,
		UserHandler:       NewUserHandler(s.UserService, l, e, secret),
		UserTaskHandler:   NewUserTaskHandler(s.UserTaskService, l, e, secret),
		TaskHandler:       NewTaskHandler(s.TaskService, l, e, secret),
		RoleHandler:       NewRoleHandler(s.RoleService, l, e, secret),
		DepartmentHandler: NewDepartmentHandler(s.DepartmentService, l, e, secret),
		CommentHandler:    NewCommentHandler(s.CommentService, l, e, secret),
		JournalHandler:    NewJournalHandler(s.JournalService, l, e, secret),
		NewsHandler:       NewNewsHandler(s.NewsService, l, e, secret),
		PrinterHandler:    NewPrinterHandler(s.PrinterService, l, e, secret),
		RoomHandler:       NewRoomHandler(s.RoomService, l, e, secret),
	}
}

var IsLoggedIn echo.MiddlewareFunc

func (h *Handler) Handle(port, secret string) error {

	IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:token",
	})

	h.UserHandler.Handle()
	h.UserTaskHandler.Handle()
	h.TaskHandler.Handle()
	h.RoleHandler.Handle()
	h.DepartmentHandler.Handle()
	h.CommentHandler.Handle()
	h.JournalHandler.Handle()
	h.NewsHandler.Handle()
	h.PrinterHandler.Handle()
	h.RoomHandler.Handle()

	err := h.e.Start(":" + port)
	if err != nil {
		return err
	}
	return nil
}
