package repository

import (
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	PrinterRepository interface {
		GetAll() ([]entity.Printer, error)
		Create(entity.PrinterCreate) error
	}

	printerRepo struct {
		db *sqlx.DB
		l  *zap.Logger
	}
)

func NewPrinterRepo(db *sqlx.DB, l *zap.Logger) PrinterRepository {
	return &printerRepo{
		db: db,
		l:  l,
	}
}

func (r printerRepo) GetAll() ([]entity.Printer, error) {
	r.l.Info("IN Printer REPO :: GET ALL Printer")

	var n []entity.Printer

	query := `
        SELECT 
            n.id, 
            n.value, 
            n.name
        FROM printers n
    `
	err := r.db.Select(&n, query)
	if err != nil {
		r.l.Error(err.Error())
		return nil, err
	}
	return n, nil
}

func (r printerRepo) Create(n entity.PrinterCreate) error {
	r.l.Debug("IN Printer SERVICE :: CREATE Printer")

	_, err := r.db.NamedExec(`INSERT INTO printers (value, name) VALUES (:value, :name)`, n)
	if err != nil {
		r.l.Error(err.Error())
		return err
	}

	return nil
}
