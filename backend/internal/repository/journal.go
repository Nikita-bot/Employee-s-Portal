package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	JournalRepository interface {
		GetJournalByTaskId(task_id int) ([]entity.TaskJournal, error)
	}
	journalRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewJournalRepo(db *sqlx.DB, l *zap.Logger) JournalRepository {
	return &journalRepo{
		db: db,
		l:  l,
	}
}

func (jr journalRepo) GetJournalByTaskId(task_id int) ([]entity.TaskJournal, error) {
	jr.l.Debug("IN REPO :: GET JOURNAL")
	var j []entity.TaskJournal

	query := `
		SELECT * FROM task_journal WHERE user_task_id=$1
	`

	err := jr.db.Select(&j, query, task_id)
	if err != nil {
		jr.l.Error(err.Error())
		return []entity.TaskJournal{}, err
	}

	return j, nil
}
