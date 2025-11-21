package handlers

import (
	"net/http"
	"portal/internal/service"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	UserHandler interface {
		Handle()
	}
	userHandler struct {
		s      service.UserService
		l      *zap.Logger
		e      *echo.Echo
		secret string
	}
)

func NewUserHandler(s service.UserService, l *zap.Logger, e *echo.Echo, secret string) UserHandler {
	return &userHandler{
		s:      s,
		l:      l,
		e:      e,
		secret: secret,
	}
}

func (uh userHandler) Handle() {
	uh.e.POST("/api/v1/auth", uh.Auth)
	uh.e.GET("/api/v1/user/:id", uh.GetUserByID, IsLoggedIn)
	uh.e.PATCH("/api/v1/user/pass", uh.PatchUser, IsLoggedIn)
	uh.e.PATCH("/api/v1/user/tg", uh.PatchUserTG, IsLoggedIn)
}

func (uh userHandler) PatchUserTG(c echo.Context) error {
	uh.l.Info("IN USER HANDLER :: CHANGE TG")
	type TGChangeRequest struct {
		ID int    `json:"user_id" form:"user_id"`
		TG string `json:"tg" form:"tg"`
	}

	var req TGChangeRequest

	req.ID, _ = strconv.Atoi(c.FormValue("user_id"))
	req.TG = c.FormValue("tg")

	uh.l.Debug("Request: ", zap.Any("req:", req))

	err := uh.s.SetTG(req.ID, req.TG)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to change tg")
	}
	return c.JSON(http.StatusOK, "Ok")
}

func (uh userHandler) PatchUser(c echo.Context) error {
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

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(uh.secret))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Expires = time.Now().Add(72 * time.Hour)
	c.SetCookie(cookie)

	request["userId"] = id

	return c.JSON(200, request)
}

func (uh userHandler) GetUserByID(c echo.Context) error {
	uh.l.Info("IN USER HANDLER :: GET USER BY ID")
	var request = make(map[string]interface{})

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	tokenUserID := claims["id"].(float64)
	tokenLogin := claims["login"].(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(400, "ID MUST BE INT")
	}

	user, err := uh.s.GetUserByID(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	if userToken == nil || float64(user.ID) != tokenUserID || user.Login != tokenLogin {
		return c.String(401, "Unauthorized")
	}

	request["user"] = user

	return c.JSON(200, request)
}
