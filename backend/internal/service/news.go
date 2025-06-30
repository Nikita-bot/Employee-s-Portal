package service

import (
	"fmt"
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	NewsService interface {
		GetAllNews() ([]entity.News, error)
		CreateNews(entity.NewsCreate) error
	}

	newsService struct {
		r repository.NewsRepository
		l *zap.Logger
	}
)

func NewNewsService(r repository.NewsRepository, l *zap.Logger) NewsService {
	return newsService{
		r: r,
		l: l,
	}
}

func (s newsService) GetAllNews() ([]entity.News, error) {
	s.l.Debug("IN SERVICE :: GET ALL NEWS")

	var n []entity.News

	n, err := s.r.GetAll()
	if err != nil {
		e := fmt.Errorf("failed to get news: %w", err)
		return nil, e
	}

	return n, nil
}

func (s newsService) CreateNews(n entity.NewsCreate) error {
	s.l.Debug("IN SERVICE :: CREATE NEWS")

	err := s.r.Create(n)
	if err != nil {
		e := fmt.Errorf("failed to create news: %w", err)
		return e
	}

	return nil
}
