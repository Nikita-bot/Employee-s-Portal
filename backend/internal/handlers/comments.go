package handlers

type (
	CommentHandler interface{}
	commentHandler struct{}
)

func NewCommentHandler() CommentHandler {
	return &commentHandler{}
}
