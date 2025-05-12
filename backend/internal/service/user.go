package service

type (
	UserService interface{}
	userService struct{}
)

func NewUserService() UserService {
	return &userService{}
}
