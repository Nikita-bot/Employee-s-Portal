package handlers

import (
	"net/http"
	"portal/internal/entity"
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	RoomHandler interface {
		Handle()
	}

	roomHandler struct {
		s service.RoomService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewRoomHandler(s service.RoomService, l *zap.Logger, e *echo.Echo) RoomHandler {
	return roomHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (h roomHandler) Handle() {
	h.e.POST("/api/v1/room", h.PostRoom)
	h.e.GET("/api/v1/room", h.GetAllRooms)
}

func (h roomHandler) PostRoom(c echo.Context) error {
	h.l.DPanic("IN HANDLER :: POST Room")

	var n entity.RoomCreate

	err := c.Bind(&n)
	if err != nil {
		h.l.Error(err.Error())
		e := "invalid parameters"
		return c.JSON(http.StatusBadRequest, e)
	}

	err = h.s.CreateRoom(n)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (h roomHandler) GetAllRooms(c echo.Context) error {
	h.l.DPanic("IN HANDLER :: GET ALL Room")

	room, err := h.s.GetAllRoom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := make(map[string]interface{})
	response["rooms"] = room

	return c.JSON(http.StatusOK, response)
}
