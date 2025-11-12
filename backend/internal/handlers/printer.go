package handlers

import (
	"net/http"
	"portal/internal/entity"
	"portal/internal/service"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	PrinterHandler interface {
		Handle()
	}

	printerHandler struct {
		s service.PrinterService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewPrinterHandler(s service.PrinterService, l *zap.Logger, e *echo.Echo) PrinterHandler {
	return printerHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (h printerHandler) Handle() {
	h.e.POST("/api/v1/printer", h.PostPrinter)
	h.e.GET("/api/v1/printer", h.GetAllPrinters)
}

func (h printerHandler) PostPrinter(c echo.Context) error {
	h.l.DPanic("IN HANDLER :: POST Printer")

	var n entity.PrinterCreate

	err := c.Bind(&n)
	if err != nil {
		h.l.Error(err.Error())
		e := "invalid parameters"
		return c.JSON(http.StatusBadRequest, e)
	}

	err = h.s.CreatePrinter(n)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (h printerHandler) GetAllPrinters(c echo.Context) error {
	h.l.DPanic("IN HANDLER :: GET ALL Printer")

	printer, err := h.s.GetAllPrinter()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := make(map[string]interface{})
	response["printers"] = printer

	return c.JSON(http.StatusOK, response)
}
