package entity

type (
	Department struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	DepartmentCreate struct {
		Name string `json:"name" db:"name"`
	}
)
