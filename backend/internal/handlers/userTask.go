package handlers

type (
	UserTaskHandler interface{}
	userTaskHandler struct{}
)

func NewUserTaskHandler() UserTaskHandler {
	return &userTaskHandler{}
}
