package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	CommentService interface{}
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
