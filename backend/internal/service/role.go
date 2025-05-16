package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	RoleService interface{}
	roleService struct {
		r repository.RoleRepository
		l *zap.Logger
	}
)

func NewRoleService(r repository.RoleRepository, l *zap.Logger) RoleService {
	return &roleService{
		r: r,
		l: l,
	}
}
