package handlers

type (
	RoleHandler interface{}
	roleHandler struct{}
)

func NewRoleHandler() RoleHandler {
	return &roleHandler{}
}
