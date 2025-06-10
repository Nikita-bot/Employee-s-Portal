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
	JournalService    JournalService
}

func NewService(l *zap.Logger, r repository.Repository) Service {
	return Service{
		UserService:       NewUserService(r.UserRepo, l),
		UserTaskService:   NewUserTaskService(r.UserTaskRepo, l),
		TaskService:       NewTaskService(r.TaskRepo, l),
		RoleService:       NewRoleService(r.RoleRepo, l),
		DepartmentService: NewDepartmentService(r.DepartmentRepo, l),
		CommentService:    NewCommentService(r.CommentRepo, l),
		JournalService:    NewJournalService(r.JournalRepo, l),
	}
}
