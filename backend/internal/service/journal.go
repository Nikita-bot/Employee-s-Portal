package service

import (
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	JournalService interface {
		GetJournalByTaskId(task_id int) ([]entity.TaskJournal, error)
	}
	journalService struct {
		r repository.JournalRepository
		l *zap.Logger
	}
)

func NewJournalService(r repository.JournalRepository, l *zap.Logger) JournalService {
	return &journalService{
		r: r,
		l: l,
	}
}

func (js journalService) GetJournalByTaskId(task_id int) ([]entity.TaskJournal, error) {
	js.l.Debug("IN SERVICE :: GET JOURNAL")

	j, err := js.r.GetJournalByTaskId(task_id)
	if err != nil {
		return []entity.TaskJournal{}, err
	}

	return j, nil
}
