package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	RoomRepository interface {
		GetAll() ([]entity.Room, error)
		Create(entity.RoomCreate) error
	}

	roomRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewRoomRepo(db *sqlx.DB, l *zap.Logger) RoomRepository {
	return &roomRepo{
		db: db,
		l:  l,
	}
}

func (r roomRepo) GetAll() ([]entity.Room, error) {
	r.l.Info("IN Room REPO :: GET ALL Room")

	var n []entity.Room

	query := `
        SELECT 
            n.id, 
            n.value, 
            n.name
        FROM rooms n
    `
	err := r.db.Select(&n, query)
	if err != nil {
		r.l.Error(err.Error())
		return nil, err
	}
	return n, nil
}

func (r roomRepo) Create(n entity.RoomCreate) error {
	r.l.Debug("IN NEWS SERVICE :: CREATE Room")

	_, err := r.db.NamedExec(`INSERT INTO rooms (value, name) VALUES (:value, :name)`, n)
	if err != nil {
		r.l.Error(err.Error())
		return err
	}

	return nil
}
