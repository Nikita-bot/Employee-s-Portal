package repository

type (
	TaskRepository interface{}

	taskRepo struct{}
)

func NewTaskRepo() TaskRepository {
	return &taskRepo{}
}
