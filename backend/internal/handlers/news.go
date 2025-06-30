package handlers

import (
	"net/http"
	"portal/internal/entity"
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	NewsHandler interface {
		Handle()
	}

	newsHandler struct {
		s service.NewsService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewNewsHandler(s service.NewsService, l *zap.Logger, e *echo.Echo) NewsHandler {
	return newsHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (h newsHandler) Handle() {
	h.e.POST("/api/v1/news", h.PostNews)
	h.e.GET("/api/v1/news", h.GetAllNews)
}

func (h newsHandler) PostNews(c echo.Context) error {
	h.l.DPanic("IN HANDLER :: POST NEWS")

	var n entity.NewsCreate

	err := c.Bind(&n)
	if err != nil {
		h.l.Error(err.Error())
		e := "invalid parameters"
		return c.JSON(http.StatusBadRequest, e)
	}

	err = h.s.CreateNews(n)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (h newsHandler) GetAllNews(c echo.Context) error {
	h.l.DPanic("IN HANDLER :: GET ALL NEWS")

	news, err := h.s.GetAllNews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := make(map[string]interface{})
	response["news"] = news

	return c.JSON(http.StatusOK, response)
}
