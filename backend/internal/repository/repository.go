package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepo       UserRepository
	UserTaskRepo   UserTaskRepository
	TaskRepo       TaskRepository
	RoleRepo       RoleRepository
	DepartmentRepo DepartmentRepository
	CommentRepo    CommetRepository
	JournalRepo    JournalRepository
}

func NewRepository(db *sqlx.DB, l *zap.Logger, r *redis.Client) Repository {
	return Repository{
		UserRepo:       NewUserRepo(db, l, r),
		UserTaskRepo:   NewUserTaskRepo(db, l, r),
		TaskRepo:       NewTaskRepo(db, l),
		RoleRepo:       NewRoleRepo(db, l),
		DepartmentRepo: NewDepartmentRepo(db, l),
		CommentRepo:    NewCommentRepo(db, l),
		JournalRepo:    NewJournalRepo(db, l),
	}
}
