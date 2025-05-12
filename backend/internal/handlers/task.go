package handlers

type (
	TaskHandler interface{}
	taskHandler struct{}
)

func NewTaskHandler() TaskHandler {
	return &taskHandler{}
}
