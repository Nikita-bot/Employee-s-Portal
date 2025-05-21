package service

import (
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	CommentService interface {
		GetComments(taskId int) ([]entity.Comment, error)
		CreateComment(cc entity.CommentCreate) error
	}
	commentService struct {
		r repository.CommetRepository
		l *zap.Logger
	}
)

func NewCommentService(r repository.CommetRepository, l *zap.Logger) CommentService {
	return &commentService{
		r: r,
		l: l,
	}
}

func (cs commentService) GetComments(taskId int) ([]entity.Comment, error) {
	cs.l.Debug("IN COMMET SERVICE :: GET ALL COMMENT FOR TASK")

	cl, err := cs.r.GetComments(taskId)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (cs commentService) CreateComment(cc entity.CommentCreate) error {
	cs.l.Debug("IN COMMET SERVICE :: CREATE NEW COMMENT")

	err := cs.r.CreateComment(cc)
	if err != nil {
		return err
	}

	return err
}
