package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	NewsRepository interface {
		GetAll() ([]entity.News, error)
		Create(entity.NewsCreate) error
	}

	newsRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewNewsRepo(db *sqlx.DB, l *zap.Logger) NewsRepository {
	return &newsRepo{
		db: db,
		l:  l,
	}
}

func (r newsRepo) GetAll() ([]entity.News, error) {
	r.l.Info("IN NEWS REPO :: GET ALL NEWS")

	var n []entity.News

	query := `
        SELECT 
            n.id, 
            n.title, 
            n.content, 
            n.date,
            u.name AS "author.name",
            u.surname AS "author.surname",
            u.patronymic AS "author.patronymic"
        FROM 
			news n
        JOIN 
			users u ON n.author=u.id 
        ORDER BY 
			n.date DESC
    `
	err := r.db.Select(&n, query)
	if err != nil {
		r.l.Error(err.Error())
		return nil, err
	}
	return n, nil
}

func (r newsRepo) Create(n entity.NewsCreate) error {
	r.l.Debug("IN USER TASK SERVICE :: CREATE TASK")

	_, err := r.db.NamedExec(`INSERT INTO news (title, author, content, date) VALUES (:title, :author, :content, :date)`, n)
	if err != nil {
		r.l.Error(err.Error())
		return err
	}

	return nil
}
