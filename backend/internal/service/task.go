package service

type (
	TaskService interface{}
	taskService struct{}
)

func NewTaskService() TaskService {
	return &taskService{}
}
