package handlers

import (
	"portal/internal/service"
	"strconv"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	DepartmentHandler interface {
		Handle()
	}
	departmentHandler struct {
		s service.DepartmentService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewDepartmentHandler(s service.DepartmentService, l *zap.Logger, e *echo.Echo) DepartmentHandler {
	return &departmentHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (dh departmentHandler) Handle() {
	dh.e.GET("/api/v1/depatments/user/:dep_id", dh.GetUsersOnDepaertmets)
}

func (dh departmentHandler) GetUsersOnDepaertmets(c echo.Context) error {
	dh.l.Info("IN DEPARTMENT HANDLER :: GET USER ON DEPARTMENT")

	var request = make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("dep_id"))
	if err != nil {
		return c.String(400, "ID MUST BE INT")
	}

	user, err := dh.s.GetUsersByDepId(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	request["users"] = user

	return c.JSON(200, request)
}
