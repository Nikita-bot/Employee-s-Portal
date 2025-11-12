package entity

type (
	Printer struct {
		ID    int    `json:"id" db:"id"`
		Value string `json:"value" db:"value"`
		Name  string `json:"name" db:"name"`
	}

	PrinterCreate struct {
		Value string `json:"value" db:"value"`
		Name  string `json:"name" db:"name"`
	}
)
