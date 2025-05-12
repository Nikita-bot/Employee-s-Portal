package service

type (
	RoleService interface{}
	roleService struct{}
)

func NewRoleService() RoleService {
	return &roleService{}
}
