package repository

type (
	CommetRepository interface{}
	commentRepo      struct{}
)

func NewCommentRepo() CommetRepository {
	return &commentRepo{}
}
