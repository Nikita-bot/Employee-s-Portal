package service

import (
	"fmt"
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	RoomService interface {
		GetAllRoom() ([]entity.Room, error)
		CreateRoom(entity.RoomCreate) error
	}

	roomService struct {
		r repository.RoomRepository
		l *zap.Logger
	}
)

func NewRoomService(r repository.RoomRepository, l *zap.Logger) RoomService {
	return roomService{
		r: r,
		l: l,
	}
}

func (s roomService) GetAllRoom() ([]entity.Room, error) {
	s.l.Debug("IN SERVICE :: GET ALL Room")

	var n []entity.Room

	n, err := s.r.GetAll()
	if err != nil {
		e := fmt.Errorf("failed to get room: %w", err)
		return nil, e
	}

	return n, nil
}

func (s roomService) CreateRoom(n entity.RoomCreate) error {
	s.l.Debug("IN SERVICE :: CREATE Room")

	err := s.r.Create(n)
	if err != nil {
		e := fmt.Errorf("failed to create room: %w", err)
		return e
	}

	return nil
}
