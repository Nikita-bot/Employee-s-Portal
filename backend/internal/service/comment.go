package service

type (
	CommentService interface{}
	commentService struct{}
)

func NewCommentService() CommentService {
	return &commentService{}
}
