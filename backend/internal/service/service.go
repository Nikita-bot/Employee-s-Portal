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
	NewsService       NewsService
	PrinterService    PrinterService
	RoomService       RoomService
}

func NewService(l *zap.Logger, r repository.Repository) Service {
	ts := NewTaskService(r.TaskRepo, l)
	return Service{
		UserService:       NewUserService(r.UserRepo, l),
		TaskService:       ts,
		RoleService:       NewRoleService(r.RoleRepo, l),
		DepartmentService: NewDepartmentService(r.DepartmentRepo, l),
		CommentService:    NewCommentService(r.CommentRepo, l),
		JournalService:    NewJournalService(r.JournalRepo, l),
		NewsService:       NewNewsService(r.NewsRepo, l),
		UserTaskService:   NewUserTaskService(r.UserTaskRepo, l, ts),
		PrinterService:    NewPrinterService(r.PrinterRepository, l),
		RoomService:       NewRoomService(r.RoomRepository, l),
	}
}
