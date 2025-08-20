package handlers

import (
	"net/http"
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
	uh.e.POST("/api/v1/auth", uh.Auth)
	uh.e.GET("/api/v1/user/:id", uh.GetUserByID)
	uh.e.PATCH("/api/v1/user/pass", uh.PatchEvent)
}

func (uh userHandler) PatchEvent(c echo.Context) error {
	uh.l.Info("IN USER HANDLER :: CHANGE PASSWORD")
	type PasswordChangeRequest struct {
		ID   int    `json:"id" form:"id"`
		Pass string `json:"pass" form:"pass"`
	}

	var req PasswordChangeRequest

	if err := c.Bind(&req); err != nil {
		uh.l.Debug("Request: ", zap.Any("req:", req))
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	uh.l.Debug("Request: ", zap.Any("req:", req))

	if req.ID == 0 || req.Pass == "" {
		return c.JSON(http.StatusBadRequest, "one of parameters is null")
	}

	err := uh.s.ChangePassword(req.ID, req.Pass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to change password")
	}
	return c.JSON(http.StatusOK, "Ok")
}

func (uh userHandler) Auth(c echo.Context) error {
	uh.l.Info("IN USER HANDLER :: AUTH")

	var request = make(map[string]interface{})

	login := c.FormValue("login")
	password := c.FormValue("password")

	id, err := uh.s.Login(login, password)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	request["userId"] = id

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
