package handlers

import (
	"portal/internal/service"
	"strconv"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	UserHandler interface {
		Handle()
	}
	userHandler struct {
		s service.UserService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewUserHandler(s service.UserService, l *zap.Logger, e *echo.Echo) UserHandler {
	return &userHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (uh userHandler) Handle() {
	uh.e.POST("/api/v1/auth", uh.getAllUser)
	uh.e.GET("/api/v1/user/:id", uh.GetUserByID)
}

func (uh userHandler) getAllUser(c echo.Context) error {
	uh.l.Info("IN USER HANDLER :: AUTH")

	var request = make(map[string]interface{})

	login := c.FormValue("login")
	password := c.FormValue("password")

	// login := c.QueryParam("login")
	// password := c.QueryParam("password")

	user, err := uh.s.Login(login, password)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	request["user"] = user

	return c.JSON(200, request)
}

func (uh userHandler) GetUserByID(c echo.Context) error {
	uh.l.Info("IN USER HANDLER :: GET USER BY ID")
	var request = make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(400, "ID MUST BE INT")
	}

	user, err := uh.s.GetUserByID(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	request["user"] = user

	return c.JSON(200, request)
}
