package entity

type (
	Task struct {
		ID           int    `json:"id" db:"id"`
		Name         string `json:"name" db:"name"`
		DepartmentID int    `json:"department_id" db:"department_id"`
		Type         string `json:"type" db:"type"`
	}

	TaskCreate struct {
		Name         string `json:"name" db:"name"`
		DepartmentID int    `json:"department_id" db:"department_id"`
	}
)
