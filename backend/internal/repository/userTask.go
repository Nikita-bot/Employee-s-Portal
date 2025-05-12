package repository

type (
	UserTaskRepository interface{}
	userTaskRepo       struct{}
)

func NewUserTaskRepo() UserTaskRepository {
	return &userTaskRepo{}
}
