package handlers

import (
	"portal/internal/service"
	"strconv"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	JournalHandler interface {
		Handle()
	}

	journalHandler struct {
		s service.JournalService
		l *zap.Logger
		e *echo.Echo
	}
)

func NewJournalHandler(s service.JournalService, l *zap.Logger, e *echo.Echo) JournalHandler {
	return &journalHandler{
		s: s,
		l: l,
		e: e,
	}
}

func (jh journalHandler) Handle() {
	jh.e.GET("/api/v1/journal/:task_id", jh.GetJournal)
}

func (jh journalHandler) GetJournal(c echo.Context) error {
	jh.l.Info("IN HANDLER :: GET JOURNAL ")
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.JSON(400, "ID MUST BE INT")
	}
	journal, err := jh.s.GetJournalByTaskId(id)
	if err != nil {
		c.JSON(500, err.Error())
	}

	response := make(map[string]interface{})
	response["journal"] = journal
	return c.JSON(200, response)
}
