package entity

type (
	Role struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	RoleCreate struct {
		Name string `json:"name" db:"name"`
	}
)
