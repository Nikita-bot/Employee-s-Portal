package handlers

import (
	"portal/internal/service"
	"strconv"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	TaskHandler interface {
		Handle()
	}
	taskHandler struct {
		s service.TaskService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewTaskHandler(s service.TaskService, l *zap.Logger, e *echo.Echo) TaskHandler {
	return &taskHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (th taskHandler) Handle() {
	th.e.GET("/api/v1/taskList", th.getAllTask)
	th.e.GET("/api/v1/tasks/support", th.getITTask)
	th.e.GET("/api/v1/tasks/users/:task_id", th.getAvailableUsers)
}

func (th taskHandler) getAvailableUsers(c echo.Context) error {
	th.l.Debug("IN TASK LIST HANDLER :: GET AVAILABLE USERS")

	var response = make(map[string]interface{})

	task_id, _ := strconv.Atoi(c.Param("task_id"))

	ul, err := th.s.GetAvailableUsers(task_id)
	if err != nil {
		return c.String(501, err.Error())
	}

	response["user_list"] = ul

	return c.JSON(200, response)
}
func (th taskHandler) getITTask(c echo.Context) error {
	th.l.Debug("IN TASK LIST HANDLER :: GET IT TASK")

	var response = make(map[string]interface{})

	tl, err := th.s.GetITTask()
	if err != nil {
		return c.String(501, err.Error())
	}

	response["task_list"] = tl

	return c.JSON(200, response)
}

func (th taskHandler) getAllTask(c echo.Context) error {
	th.l.Debug("IN TASK LIST HANDLER :: GET ALL TASK")

	var response = make(map[string]interface{})

	tl, err := th.s.GetAllTask()
	if err != nil {
		return c.String(501, err.Error())
	}

	response["task_list"] = tl

	return c.JSON(200, response)
}
