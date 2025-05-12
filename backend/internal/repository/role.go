package repository

type (
	RoleRepository interface{}
	roleRepo       struct{}
)

func NewRoleRepo() RoleRepository {
	return &roleRepo{}
}
