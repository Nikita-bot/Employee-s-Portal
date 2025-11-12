package entity

type (
	Task struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
		Type string `json:"type" db:"type"`
	}

	TaskCreate struct {
		Name string `json:"name" db:"name"`
		Type string `json:"type" db:"type"`
	}
)
