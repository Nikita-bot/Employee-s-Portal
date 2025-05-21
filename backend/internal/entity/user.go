package entity

type (
	User struct {
		ID             int        `json:"id" db:"id"`
		Name           string     `json:"name" db:"name"`
		Surname        string     `json:"surname" db:"surname"`
		Patronymic     string     `json:"patronymic" db:"patronymic"`
		Position       string     `json:"position" db:"position"`
		Roles          []Role     `json:"roles"`
		Department     Department `json:"department"`
		Email          string     `json:"email" db:"email"`
		Phone          string     `json:"phone" db:"phone"`
		TgLink         string     `json:"tg_link" db:"tg_link"`
		TgID           string     `json:"tg_id" db:"tg_id"`
		Pasport        string     `json:"pasport" db:"pasport"`
		Snyls          string     `json:"snyls" db:"snyls"`
		Adress         string     `json:"adress" db:"adress"`
		EmploymentDate string     `json:"employment_date" db:"employment_date"`
		DissmissalDate string     `json:"dismissal_date" db:"dismissal_date"`
	}

	UserCreate struct {
		Name           string `json:"name" db:"name"`
		Surname        string `json:"surname" db:"surname"`
		Patronymic     string `json:"patronymic" db:"patronymic"`
		Position       string `json:"position" db:"position"`
		DepartmentID   int    `json:"department_id" db:"department_id"`
		Login          string `json:"login" db:"login"`
		Pass           string `json:"password" db:"password"`
		Email          string `json:"email" db:"email"`
		Phone          string `json:"phone" db:"phone"`
		TgLink         string `json:"tg_link" db:"tg_link"`
		TgID           string `json:"tg_id" db:"tg_id"`
		Pasport        string `json:"pasport" db:"pasport"`
		Snyls          string `json:"snyls" db:"snyls"`
		Adress         string `json:"adress" db:"adress"`
		EmploymentDate string `json:"employment_date" db:"employment_date"`
		DissmissalDate string `json:"dismissal_date" db:"dismissal_date"`
	}
)
