package entity

type (
	User struct {
		ID         int      `json:"id" db:"id"`
		Surname    string   `json:"surname" db:"surname"`
		Name       string   `json:"name" db:"name"`
		Patronymic string   `json:"patronymic" db:"patronymic"`
		Login      string   `json:"login" db:"login"`
		Password   string   `json:"password" db:"password"`
		Roles      []Role   `json:"roles"`
		Snyls      string   `json:"snyls" db:"snyls"`
		PassSer    string   `json:"pasport_ser" db:"pasport_ser"`
		PassNum    string   `json:"pasport_num" db:"pasport_num"`
		PassDate   string   `json:"pasport_date" db:"pasport_date"`
		PassWho    string   `json:"pasport_dep" db:"pasport_dep"`
		PassWhoKey string   `json:"pasport_dep_key" db:"pasport_dep_key"`
		Adress     string   `json:"adress" db:"adress"`
		Phone      string   `json:"phone" db:"phone"`
		Email      string   `json:"email" db:"email"`
		TgLink     string   `json:"tg_link" db:"tg_link"`
		TgID       string   `json:"tg_id" db:"tg_id"`
		Employee   Employee `json:"employee" db:"employee"`
	}

	UserMainData struct {
		ID         int    `json:"id" db:"id"`
		Name       string `json:"name" db:"name"`
		Surname    string `json:"surname" db:"surname"`
		Patronymic string `json:"patronymic" db:"patronymic"`
		Department string `json:"department" db:"department"`
	}
	// ПОПРАВИТЬ BRANCH_ID В РЕПОЗИТОРИИ И НА КЛИЕНТЕ
	Employee struct {
		ID            string       `json:"id" db:"id"`
		TabNum        string       `json:"tab_num" db:"tab_num"`
		UserID        int          `json:"user_id" db:"user_id"`
		Zanyatost     string       `json:"zanyatost" db:"zanyatost"`
		StartDate     string       `json:"start_date" db:"start_date"`
		EndDate       string       `json:"end_date" db:"end_date"`
		Position      string       `json:"position" db:"position"`
		OtherPosition []Position   `json:"other_position"`
		DepartmentID  string       `db:"depart_id"`
		Department    string       `json:"department"`
		BranchID      int          `json:"branch_id" db:"branch_id"`
		Branch        string       `json:"branch"`
		BossID        int          `db:"boss"`
		Boss          UserMainData `json:"boss"`
	}

	Position struct {
		Name       string `json:"name" db:"name"`
		Department string `json:"department" db:"department"`
	}
)
