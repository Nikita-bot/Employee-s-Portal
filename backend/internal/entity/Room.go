package entity

type (
	Room struct {
		ID    int    `json:"id"`
		Value string `json:"value" db:"value"`
		Name  string `json:"name"`
	}

	RoomCreate struct {
		Value string `json:"value" db:"value"`
		Name  string `json:"name"`
	}
)
