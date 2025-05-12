package service

type (
	DepartmentService interface{}
	departmentService struct{}
)

func NewDepartmentService() DepartmentService {
	return &departmentService{}
}
