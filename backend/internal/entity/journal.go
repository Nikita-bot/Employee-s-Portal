package entity

type (
	TaskJournal struct {
		UserTaskID    int    `json:"user_task_id" form:"user_task_id" db:"user_task_id"`
		Action        string `json:"action" form:"action" db:"action"`
		Creation_date string `json:"creation_date" form:"creation_date" db:"creation_date"`
	}
)
