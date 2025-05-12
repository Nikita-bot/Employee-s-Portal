package repository

type (
	UserRepository interface {
	}

	userRepo struct{}
)

func NewUserRepo() UserTaskRepository {
	return &userRepo{}
}
