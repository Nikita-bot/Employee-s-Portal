package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepo       UserRepository
	UserTaskRepo   UserTaskRepository
	TaskRepo       TaskRepository
	RoleRepo       RoleRepository
	DepartmentRepo DepartmentRepository
	CommentRepo    CommetRepository
}

func NewRepository(db *sqlx.DB, l *zap.Logger) Repository {
	return Repository{
		UserRepo:       NewUserRepo(),
		UserTaskRepo:   NewUserTaskRepo(),
		TaskRepo:       NewTaskRepo(),
		RoleRepo:       NewRoleRepo(),
		DepartmentRepo: NewDepartmentRepo(),
		CommentRepo:    NewCommentRepo(),
	}
}
