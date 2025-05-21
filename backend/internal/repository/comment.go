package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	CommetRepository interface {
		GetComments(taskId int) ([]entity.Comment, error)
		CreateComment(cc entity.CommentCreate) error
	}
	commentRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewCommentRepo(db *sqlx.DB, l *zap.Logger) CommetRepository {
	return &commentRepo{
		db: db,
		l:  l,
	}
}

func (cr commentRepo) GetComments(taskId int) ([]entity.Comment, error) {
	cr.l.Info("IN COMMENT REPO :: GET ALL COMMENTS FOR TASK " + string(taskId))

	var c []entity.Comment

	query := `
		SELECT * FROM comments WHERE user_task_id=$1
	`

	err := cr.db.Select(&c, query, taskId)
	if err != nil {
		cr.l.Error(err.Error())
		return nil, err
	}

	return c, nil
}

func (cr commentRepo) CreateComment(cc entity.CommentCreate) error {
	cr.l.Info("IN COMMET REPO :: CREATE NEW COMMET")

	query := `
		INCERT INTO comments (user_task_id, author_id, comment, creation_date)
		VALUES (:user_task_id, :author_id, :comment, :creation_date)
	`

	_, err := cr.db.NamedExec(query, cc)
	if err != nil {
		cr.l.Error(err.Error())
		return err
	}

	return nil
}
