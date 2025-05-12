package handlers

type (
	UserHandler interface{}
	userHandler struct{}
)

func NewUserHandler() UserHandler {
	return &userHandler{}
}
