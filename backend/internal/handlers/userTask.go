package handlers

import (
	"net/http"
	"portal/internal/entity"
	"portal/internal/service"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	UserTaskHandler interface {
		Handle()
	}
	userTaskHandler struct {
		s service.UserTaskService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewUserTaskHandler(s service.UserTaskService, l *zap.Logger, e *echo.Echo) UserTaskHandler {
	return &userTaskHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (uh userTaskHandler) Handle() {
	uh.e.GET("/api/v1/tasks/user/:userId", uh.getAllUserTask)        //Получить все задачи пользователя
	uh.e.GET("/api/v1/tasks/:id", uh.getUserTaskByID)                //Получить задачу по id
	uh.e.POST("/api/v1/tasks", uh.postUserTask)                      //Создать задачу для пользователя
	uh.e.PATCH("/api/v1/tasks/:id", uh.patchUserTask)                //Выполнить задачу
	uh.e.PATCH("/api/v1/tasks/executor/:id", uh.patchExecutorToTask) //Сменить исполнителя задачи
	uh.e.DELETE("/api/v1/tasks/:id", uh.deleteUserTask)              //Удалить задачу
}

func (uh userTaskHandler) patchExecutorToTask(c echo.Context) error {
	uh.l.Info("IH USER TASK HANDLER :: UPDATE TASK")

	type ChangeExecutorRequest struct {
		ExecutorID int `json:"executor_id" validate:"required"`
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(400, "Task ID is required")
	}

	var req ChangeExecutorRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request body"})
	}

	err = uh.s.ChangeExecutor(taskID, req.ExecutorID)
	if err != nil {
		return c.String(500, "Can`t change executor")
	}

	return c.String(200, "Task updated successfully")
}

func (uh userTaskHandler) getAllUserTask(c echo.Context) error {
	uh.l.Info("IH USER TASK HANDLER :: GET ALL TASK FOR USER")

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.String(401, "ID must be int")
	}

	taskByUser, err := uh.s.TaskByUser(userId)
	if err != nil {
		return c.String(501, err.Error())
	}

	taskForUser, err := uh.s.TaskForUser(userId)
	if err != nil {
		return c.String(501, err.Error())
	}

	var response = make(map[string]interface{})
	response["initiator"] = taskByUser
	response["executor"] = taskForUser

	return c.JSON(200, response)
}

func (uh userTaskHandler) postUserTask(c echo.Context) error {
	uh.l.Info("IH USER TASK HANDLER :: CREATE TASK FOR USER")

	var uc entity.UserTaskCreate

	err := c.Bind(&uc)
	if err != nil {
		return c.String(401, "bad request")
	}

	uh.l.Debug("IN USER TASK SERVICE", zap.Int("branch_id", uc.BranchID))

	err = uh.s.CreateTask(uc)
	if err != nil {
		return c.String(501, err.Error())
	}

	return c.String(200, "Ok")
}

func (uh userTaskHandler) getUserTaskByID(c echo.Context) error {
	uh.l.Info("IH USER TASK HANDLER :: GET TASK BY ID")

	var response = make(map[string]interface{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(401, "ID must be int")
	}

	ut, err := uh.s.GetTaskByID(id)
	if err != nil {
		return c.String(501, err.Error())
	}

	response["task"] = ut

	return c.JSON(200, response)
}

func (uh *userTaskHandler) deleteUserTask(c echo.Context) error {
	uh.l.Info("IH USER TASK HANDLER :: DELETE TASK")

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(400, "Task ID is required")
	}

	err = uh.s.DeleteTask(taskID)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.String(200, "Task deleted successfully")
}

func (uh *userTaskHandler) patchUserTask(c echo.Context) error {
	uh.l.Info("IH USER TASK HANDLER :: UPDATE TASK")

	resp := struct {
		ID     int `json:"id"`
		Status int `json:"status" form:"status"`
	}{}

	// Получаем ID задачи из URL
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(400, "Task ID is required")
	}

	err = c.Bind(&resp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad status is required")
	}

	execDate := ""

	if resp.Status == 2 {
		now := time.Now()
		execDate = now.Format("2006-01-02")
	}

	err = uh.s.ExecTask(taskID, resp.Status, execDate)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.String(200, "Task updated successfully")
}
