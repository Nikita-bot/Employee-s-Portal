package entity

type (
	UserTask struct {
		ID          int          `json:"id" db:"id"`
		Task        string       `json:"task_id" db:"task_id"`
		Executor    UserMainData `json:"executor" db:"executor"`
		Initiator   UserMainData `json:"initiator" db:"initiator"`
		Description string       `json:"description" db:"description"`
		Status      int          `json:"status" db:"status"`
		CreateDate  string       `json:"create_date" db:"create_date"`
		ExecuteDate string       `json:"execute_date" db:"execute_date"`
	}

	UserTaskCreate struct {
		Task        int    `json:"task_id" form:"task_id" db:"task_id"`
		Executor    int    `json:"executor" form:"executor" db:"executor"`
		Initiator   int    `json:"initiator" form:"initiator" db:"initiator"`
		Description string `json:"description" form:"description" db:"description"`
		Status      int    `json:"status" form:"status" db:"status"`
		CreateDate  string `json:"create_date" form:"create_date" db:"create_date"`
	}

	UserCountTask struct {
		UserID int `db:"user_id"`
		Count  int `db:"task_count"`
	}
)
