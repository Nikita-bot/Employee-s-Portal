package service

import (
	"portal/internal/repository"

	"go.uber.org/zap"
)

type Service struct {
	UserService       UserService
	UserTaskService   UserTaskService
	TaskService       TaskService
	RoleService       RoleService
	DepartmentService DepartmentService
	CommentService    CommentService
}

func NewService(l *zap.Logger, r repository.Repository) Service {
	return Service{
		UserService:       NewUserService(),
		UserTaskService:   NewUserTaskService(),
		TaskService:       NewTaskService(),
		RoleService:       NewRoleService(),
		DepartmentService: NewDepartmentService(),
		CommentService:    NewCommentService(),
	}
}
