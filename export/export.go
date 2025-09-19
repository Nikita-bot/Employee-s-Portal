package main

import (
	"encoding/json"
	"export/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Employee struct {
	ID        string `json:"УникальныйИдентификатор" db:"id"`
	TabNum    string `json:"ТабельныйНомер" db:"tab_num"`
	UserData  User   `json:"ФизическоеЛицо"`
	UserID    int    `db:"user_id"`
	Zanyatost string `json:"ВидЗанятости" db:"zanyatost"`
	StartDate string `json:"ДатаПриема" db:"start_date"`
	EndDate   string `json:"ДатаУвольнения" db:"end_date"`
	Dol       Dolj   `json:"Должность"`
	Dol_Name  string `db:"position"`
	Boss      int    `db:"boss"`
	Depart    Depart `json:"Подразделение"`
	DepartID  int    `db:"depart_id"`
	BranchID  int    `db:"branch_id"`
}

type User struct {
	ID         string `json:"УникальныйИдентификатор"`
	Fam        string `json:"Фамилия" db:"surname"`
	Name       string `json:"Имя" db:"name"`
	Otc        string `json:"Отчество" db:"patronymic"`
	Snyls      string `json:"СтраховойНомерПФР" db:"snyls"`
	Login      string `db:"login"`
	PassSer    string `json:"ДокументСерия" db:"pasport_ser"`
	PassNum    string `json:"ДокументНомер" db:"pasport_num"`
	PassDate   string `json:"ДокументДатаВыдачи" db:"pasport_date"`
	PassWho    string `json:"ДокументКемВыдан" db:"pasport_dep"`
	PassWhoKey string `json:"ДокументКодПодразделения" db:"pasport_dep_key"`
	Adress     string `json:"АдресПоПропискеПредставление" db:"adress"`
	Number     string `json:"ТелефонМобильныйПредставление" db:"phone"`
	Emain      string `json:"EMailПредставление" db:"email"`
}

type Dolj struct {
	ID   string `json:"УникальныйИдентификатор"`
	Name string `json:"Наименование" db:"name"`
}

type Depart struct {
	ID      string `json:"УникальныйИдентификатор"`
	Name    string `json:"Наименование" db:"name"`
	IsAlive string `json:"Расформировано"`
	IsDie   bool   `db:"is_die"`
}

var depInBranch2 = map[string]string{
	"аптечный пункт":                         "",
	"гинекологическое отделение №2":          "",
	"детская поликлиника":                    "",
	"дневной стационар":                      "",
	"женская консультация":                   "",
	"Клинико-диагностическая лаборатория №2": "",
	"неврологическое отделение №2":           "",
	"ожоговое отделение, в том числе палата реанимации и интенсивной терапии": "",
	"отделение анестезиологии-реанимации №2":                                  "",
	"отделение анестезиологии-реанимации №4":                                  "",
	"отделение медицинской профилактики":                                      "",
	"Отделение медицинской реабилитации для взрослых":                         "",
	"отделение новорожденных №1":                                              "",
	"Отделение общей врачебной практики №1":                                   "",
	"Отделение общей врачебной практики №2":                                   "",
	"отделение травматологии и ортопедии №3":                                  "",
	"отделение ультразвуковой диагностики №2":                                 "",
	"педиатрическое отделение":                                                "",
	"педиатрическое отделение №1":                                             "",
	"педиатрическое отделение №2":                                             "",
	"первичное сосудистое отделение":                                          "",
	"поликлиника":                     "",
	"приемное отделение №2":           "",
	"рентгенологическое отделение №2": "",
	"терапевтическое отделение":       "",
	"терапевтическое отделение №1, в том числе палата интенсивной терапии": "",
	"травматолого-ортопедическое отделение":                                "",
	"урологическое отделение №2":                                           "",
	"физиотерапевтическое отделение №2":                                    "",
	"хирургическое отделение №3":                                           "",
	"эндоскопическое отделение №2":                                         "",
}

func main() {

	conf := config.New()
	db := InitDB(*conf)
	defer db.Close()

	processEmployees(conf, db)

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		processEmployees(conf, db)
	}
}

func processEmployees(conf *config.Config, db *sqlx.DB) {
	var url string

	log.Println("Начало получения данных о новых сотрудниках")

	isUsersExist, err := CheckUsersInDB(db)
	if err != nil {
		log.Println("Не получилось получить данные о пользователях: ", err.Error())
		return
	}

	if isUsersExist {
		now := time.Now()
		dateStr := now.Format("20060102")
		url = fmt.Sprintf(conf.API_URL, dateStr)
	} else {
		url = conf.FIRST_API_URL
	}

	log.Println(url)

	employees := GetEmployees(url)
	if len(employees) == 0 {
		log.Println("Нет новых данных о сотрудниках")
		return
	}
	log.Println("Получены данные о сотрудниках")

	for i, emp := range employees {
		var err error

		// Проверяем, есть ли сотрудник уже в базе
		err = GetEmployeesFromBD(emp.ID, db)
		if err == nil {
			log.Println("Сотрудник " + emp.ID + " уже в базе")
			continue
		}

		employees[i].Dol_Name = emp.Dol.Name

		employees[i].DepartID, err = GetDepartmetsByName(emp.Depart.Name, db)
		if err != nil {
			id, err := CreateDepartmets(emp.Depart, db)
			if err != nil {
				log.Println("Failed to create Department: ", err.Error())
				continue
			}
			employees[i].DepartID = id
		}

		// Получаем или создаем пользователя
		employees[i].UserID, err = GetUserBySnyls(emp.UserData.Snyls, db)
		if err != nil {
			user_id, err := CreateUser(emp.UserData, db)
			if err != nil {
				log.Println("Failed to create User: ", err.Error())
				continue
			}
			employees[i].UserID = user_id
		}

		_, exists := depInBranch2[emp.Depart.Name]
		if exists {
			employees[i].BranchID = 2
		} else {
			employees[i].BranchID = 1
		}

		// Создаем сотрудника
		err = CreateEmployee(employees[i], db)
		if err != nil {
			log.Println("Failed to create employee: ", err.Error())
		}
	}
	log.Println("Все новые сотрудники записаны")
}

func CheckUsersInDB(db *sqlx.DB) (bool, error) {
	// log.Println("Check users")

	var count int

	query := `
		SELECT COUNT(*) FROM users
	`

	err := db.Get(&count, query)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func CreateDepartmets(d Depart, db *sqlx.DB) (int, error) {
	log.Println("CREATE DEPARTMENT")

	var id int64

	query := `
		INSERT INTO departments (name, is_die) VALUES (:name, :is_die) RETURNING id
	`

	if d.IsAlive == "Нет" {
		d.IsDie = false
	} else if d.IsAlive == "Да" {
		d.IsDie = true
	}

	rows, err := db.NamedQuery(query, d)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalln("failed to get id: ", err.Error())
			return 0, err
		}
	}

	return int(id), err
}

func GetDepartmetsByName(name string, db *sqlx.DB) (int, error) {
	// log.Println("Get Departments")

	var id int

	query := `
		SELECT id FROM departments WHERE name=$1;
	`

	err := db.Get(&id, query, name)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func CreateUser(u User, db *sqlx.DB) (int, error) {
	// log.Println("CREATE USER")

	var id int64

	query := `
		INSERT INTO users (surname, name, patronymic, snyls, pasport_ser, pasport_num, pasport_date, pasport_dep, pasport_dep_key, adress, phone, email, login) 
		VALUES (:surname, :name, :patronymic, :snyls, :pasport_ser, :pasport_num, :pasport_date, :pasport_dep, :pasport_dep_key, :adress, :phone, :email, :login)
		RETURNING id
		`

	u.Login = fmt.Sprintf("%s %s %s", u.Fam, u.Name, u.Otc)

	rows, err := db.NamedQuery(query, u)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalln("failed to get id: ", err.Error())
			return 0, err
		}
	}

	return int(id), err
}

func GetUserBySnyls(snyls string, db *sqlx.DB) (int, error) {
	// log.Println("Get User")

	var id int

	query := `
		SELECT id FROM users WHERE snyls=$1
	`

	err := db.Get(&id, query, snyls)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func CreateEmployee(e Employee, db *sqlx.DB) error {

	query := `
		INSERT INTO employee (id, tab_num, user_id, zanyatost, start_date, end_date, position, depart_id, branch_id) 
		VALUES (:id, :tab_num, :user_id, :zanyatost, :start_date, :end_date, :position, :depart_id, :branch_id)
	`

	_, err := db.NamedExec(query, e)
	if err != nil {
		return err
	}

	return nil
}

func GetEmployeesFromBD(id string, db *sqlx.DB) error {
	// log.Println("Get User")

	var e Employee

	query := `
		SELECT * FROM employee WHERE id=$1
	`

	err := db.Get(&e, query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetEmployees(apiUrl string) []Employee {
	// log.Println("Get Employees")
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Fatalf("Ошибка при создании запроса: %v", err)
	}

	req.SetBasicAuth("apiuser", "pfv,tpbz")

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	var employees []Employee
	err = json.Unmarshal(body, &employees)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	return employees
}

func InitDB(c config.Config) *sqlx.DB {
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Pg_host, c.Pg_port, c.Pg_user, c.Db_name, c.Pg_pass)
	db, err := sqlx.Open("postgres", str)
	if err != nil {
		log.Panic("Ошибка подключения к БД: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Panic("Ошибка подключения к БД: " + err.Error())
	}

	return db
}
