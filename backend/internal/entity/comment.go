package entity

type (
	Comment struct {
		ID            int          `json:"id" db:"id"`
		UserTaskId    int          `json:"user_task_id" db:"user_task_id"`
		Author        UserMainData `json:"author" db:"author"`
		Comment       string       `json:"comment" db:"comment"`
		Creation_date string       `json:"creation_date" db:"creation_date"`
	}

	CommentCreate struct {
		UserTaskId    int    `json:"user_task_id" form:"user_task_id" db:"user_task_id"`
		Author        int    `json:"author_id" form:"author_id" db:"author_id"`
		Comment       string `json:"comment" form:"comment" db:"comment"`
		Creation_date string `json:"creation_date" form:"creation_date" db:"creation_date"`
	}
)
