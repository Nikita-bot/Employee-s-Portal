package handlers

type (
	DepartmentHandler interface{}
	departmentHandler struct{}
)

func NewDepartmentHandler() DepartmentHandler {
	return &departmentHandler{}
}
