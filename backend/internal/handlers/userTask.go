package handlers

import (
	"portal/internal/entity"
	"portal/internal/service"
	"strconv"

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
	uh.e.GET("/api/v1/tasks/user/:userId", uh.getAllUserTask)
	uh.e.POST("/api/v1/tasks", uh.postUserTask)
	uh.e.GET("/api/v1/tasks/:id", uh.getUserTaskByID)
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
