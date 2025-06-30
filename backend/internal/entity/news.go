package entity

type (
	News struct {
		ID      int    `json:"id" db:"id"`
		Title   string `json:"title" db:"title"`
		Author  User   `json:"author" db:"author"`
		Content string `json:"content" db:"content"`
		Date    string `json:"date" db:"date"`
	}

	NewsCreate struct {
		Title   string `json:"title" db:"title"`
		Author  int    `json:"author" db:"author"`
		Content string `json:"content" db:"content"`
		Date    string `json:"date" db:"date"`
	}
)
