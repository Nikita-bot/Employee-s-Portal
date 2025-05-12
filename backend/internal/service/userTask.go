package service

type (
	UserTaskService interface{}
	userTaskService struct{}
)

func NewUserTaskService() UserTaskService {
	return &userTaskService{}
}
