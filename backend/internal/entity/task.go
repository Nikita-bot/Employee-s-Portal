package entity

type (
	Task struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	TaskCreate struct {
		Name string `json:"name" db:"name"`
	}
)
