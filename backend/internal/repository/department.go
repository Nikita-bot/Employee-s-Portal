package repository

type (
	DepartmentRepository interface{}
	departmentRepo       struct{}
)

func NewDepartmentRepo() DepartmentRepository {
	return &departmentRepo{}
}
