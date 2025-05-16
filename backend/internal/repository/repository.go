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
		UserRepo:       NewUserRepo(db, l),
		UserTaskRepo:   NewUserTaskRepo(db, l),
		TaskRepo:       NewTaskRepo(db, l),
		RoleRepo:       NewRoleRepo(db, l),
		DepartmentRepo: NewDepartmentRepo(db, l),
		CommentRepo:    NewCommentRepo(db, l),
	}
}
