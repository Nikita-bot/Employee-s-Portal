package service

import (
	"fmt"
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	PrinterService interface {
		GetAllPrinter() ([]entity.Printer, error)
		CreatePrinter(entity.PrinterCreate) error
	}

	printerService struct {
		r repository.PrinterRepository
		l *zap.Logger
	}
)

func NewPrinterService(r repository.PrinterRepository, l *zap.Logger) PrinterService {
	return printerService{
		r: r,
		l: l,
	}
}

func (s printerService) GetAllPrinter() ([]entity.Printer, error) {
	s.l.Debug("IN SERVICE :: GET ALL Printer")

	var n []entity.Printer

	n, err := s.r.GetAll()
	if err != nil {
		e := fmt.Errorf("failed to get printer: %w", err)
		return nil, e
	}

	return n, nil
}

func (s printerService) CreatePrinter(n entity.PrinterCreate) error {
	s.l.Debug("IN SERVICE :: CREATE Printer")

	err := s.r.Create(n)
	if err != nil {
		e := fmt.Errorf("failed to create printer: %w", err)
		return e
	}

	return nil
}
