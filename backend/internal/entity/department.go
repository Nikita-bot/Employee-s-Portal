package entity

type (
	Department struct {
		ID     int    `json:"id" db:"id"`
		Name   string `json:"name" db:"name"`
		IsDied string `json:"is_die" db:"is_die"`
	}

	DepartmentCreate struct {
		Name   string `json:"name" db:"name"`
		IsDied string `json:"is_die" db:"is_die"`
	}
)
