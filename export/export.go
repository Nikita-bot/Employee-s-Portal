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

type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
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

const (
	RoleAdministration    = "Администрация"
	RoleMedicalManagement = "Медицинские руководители"
	RoleDoctors           = "Врачи"
	RoleNursingStaff      = "Средний медицинский персонал"
	RoleJuniorMedical     = "Младший медицинский персонал"
	RoleTechnical         = "Технический персонал"
	RoleEconomic          = "Экономика и бухгалтерия"
	RoleIT                = "IT-специалисты"
	RoleHR                = "Кадры и административный персонал"
	RoleSupport           = "Обслуживающий персонал"
)

// RoleMapping сопоставляет должности с ролями
var RoleMapping = map[string]string{
	// АДМИНИСТРАЦИЯ
	"Главный врач": RoleAdministration,
	"Заместитель главного врача по акушерской и гинекологической помощи":                        RoleAdministration,
	"Заместитель главного врача по амбулаторно-поликлинической помощи":                          RoleAdministration,
	"Заместитель главного врача по клинико-диагностической службе":                              RoleAdministration,
	"Заместитель главного врача по клинико-экспертной работе":                                   RoleAdministration,
	"Заместитель главного врача по организационно-методической работе":                          RoleAdministration,
	"Заместитель главного врача по терапии":                                                     RoleAdministration,
	"Заместитель главного врача по хирургической помощи":                                        RoleAdministration,
	"Заместитель главного врача по экономическим вопросам":                                      RoleAdministration,
	"Начальник отдела по борьбе с коррупцией":                                                   RoleAdministration,
	"Начальник организационно-методического отдела":                                             RoleAdministration,
	"Начальник отдела автоматизированных систем управления":                                     RoleAdministration,
	"Начальник отдела документационного обеспечения":                                            RoleAdministration,
	"Начальник отдела информации и общественных связей":                                         RoleAdministration,
	"Начальник отдела материально-технического снабжения":                                       RoleAdministration,
	"Начальник отдела по внутреннему контролю качества и безопасности медицинской деятельности": RoleAdministration,
	"Начальник отдела по обращению с отходами":                                                  RoleAdministration,
	"Начальник отдела ресурсного обеспечения":                                                   RoleAdministration,
	"Начальник отдела статистики":                                                               RoleAdministration,
	"Начальник отдела экономического анализа":                                                   RoleAdministration,
	"Начальник службы охраны труда":                                                             RoleAdministration,
	"Начальник службы психологической помощи":                                                   RoleAdministration,
	"Начальник юридического отдела":                                                             RoleAdministration,

	// МЕДИЦИНСКИЕ РУКОВОДИТЕЛИ (ЗАВЕДУЮЩИЕ ОТДЕЛЕНИЯМИ)
	"Заведующий аптечным пунктом - провизор":                                                                    RoleMedicalManagement,
	"Заведующий бактериологической лабораторией - врач-бактериолог":                                             RoleMedicalManagement,
	"Заведующий гастроэнтерологическим отделением №1 - врач-гастроэнтеролог":                                    RoleMedicalManagement,
	"Заведующий гинекологическим отделением №1 - врач-акушер-гинеколог":                                         RoleMedicalManagement,
	"Заведующий гинекологическим отделением №2 - врач-акушер-гинеколог":                                         RoleMedicalManagement,
	"Заведующий детской поликлиникой - врач-педиатр":                                                            RoleMedicalManagement,
	"Заведующий женской консультацией - врач-акушер-гинеколог":                                                  RoleMedicalManagement,
	"Заведующий кардиологическим отделением №1 - врач-кардиолог":                                                RoleMedicalManagement,
	"Заведующий клинико-диагностической лабораторией №1 - врач клинической лабораторной диагностики":            RoleMedicalManagement,
	"Заведующий клинико-диагностической лабораторией №2 - врач клинической лабораторной диагностики":            RoleMedicalManagement,
	"Заведующий консультативной поликлиникой - врач-терапевт":                                                   RoleMedicalManagement,
	"Заведующий неврологическим отделением №1 - врач-невролог":                                                  RoleMedicalManagement,
	"Заведующий неврологическим отделением №2 - врач-невролог":                                                  RoleMedicalManagement,
	"Заведующий нейрохирургическим отделением - врач-нейрохирург":                                               RoleMedicalManagement,
	"Заведующий ожоговым отделением - врач-хирург":                                                              RoleMedicalManagement,
	"Заведующий операционным блоком №1 - врач-травматолог-ортопед":                                              RoleMedicalManagement,
	"Заведующий отделением анестезиологии-реанимации №1 - врач-анестезиолог-реаниматолог":                       RoleMedicalManagement,
	"Заведующий отделением анестезиологии-реанимации №2 - врач-анестезиолог-реаниматолог":                       RoleMedicalManagement,
	"Заведующий отделением анестезиологии-реанимации №4 - врач-анестезиолог-реаниматолог":                       RoleMedicalManagement,
	"Заведующий отделением медицинской профилактики - врач по медицинской профилактике":                         RoleMedicalManagement,
	"Заведующий отделением общей врачебной практики №2 - врач общей практики (семейный врач)":                   RoleMedicalManagement,
	"Заведующий отделением острых отравлений - врач-токсиколог":                                                 RoleMedicalManagement,
	"Заведующий отделением переливания крови - врач-трансфузиолог":                                              RoleMedicalManagement,
	"Заведующий отделением травматологии и ортопедии №1 - врач-травматолог-ортопед":                             RoleMedicalManagement,
	"Заведующий отделением травматологии и ортопедии №2 - врач-травматолог-ортопед":                             RoleMedicalManagement,
	"Заведующий отделением травматологии и ортопедии №3 - врач-травматолог-ортопед":                             RoleMedicalManagement,
	"Заведующий отделением ультразвуковой диагностики №1 - врач ультразвуковой диагностики":                     RoleMedicalManagement,
	"Заведующий отделением ультразвуковой диагностики №2 - врач ультразвуковой диагностики":                     RoleMedicalManagement,
	"Заведующий отделением функциональной диагностики №1 - врач функциональной диагностики":                     RoleMedicalManagement,
	"Заведующий отделением экспертизы временной нетрудоспособности - врач-методист":                             RoleMedicalManagement,
	"Заведующий отделением эндокринологии - врач-эндокринолог":                                                  RoleMedicalManagement,
	"Заведующий отделом внебюджетной деятельности - врач-методист":                                              RoleMedicalManagement,
	"Заведующий отделом телемедицинских консультаций - врач-хирург":                                             RoleMedicalManagement,
	"Заведующий отделом экспертизы и взаимодействия со страховыми медицинскими организациями - врач-неонатолог": RoleMedicalManagement,
	"Заведующий оториноларингологическим отделением - врач-оториноларинголог":                                   RoleMedicalManagement,
	"Заведующий педиатрическим отделением - врач-педиатр":                                                       RoleMedicalManagement,
	"Заведующий педиатрическим отделением №1 - врач педиатр участковый":                                         RoleMedicalManagement,
	"Заведующий педиатрическим отделением №2 - врач-педиатр":                                                    RoleMedicalManagement,
	"Заведующий первичным сосудистым отделением - врач-невролог":                                                RoleMedicalManagement,
	"Заведующий поликлиникой - врач-кардиолог":                                                                  RoleMedicalManagement,
	"Заведующий приемным отделением №1 - врач-кардиолог":                                                        RoleMedicalManagement,
	"Заведующий приемным отделением №2 - врач-невролог":                                                         RoleMedicalManagement,
	"Заведующий пульмонологическим отделением №1 - врач-пульмонолог":                                            RoleMedicalManagement,
	"Заведующий терапевтическим отделением - врач-терапевт участковый":                                          RoleMedicalManagement,
	"Заведующий терапевтическим отделением №1 - врач-терапевт":                                                  RoleMedicalManagement,
	"Заведующий травматолого-ортопедическим отделением - врач-травматолог-ортопед":                              RoleMedicalManagement,
	"Заведующий урологическим отделением №1 - врач-уролог":                                                      RoleMedicalManagement,
	"Заведующий урологическим отделением №2 - врач-уролог":                                                      RoleMedicalManagement,
	"Заведующий учебно-методическим отделом":                                                                    RoleMedicalManagement,
	"Заведующий физиотерапевтическим отделением №1 - врач-физиотерапевт":                                        RoleMedicalManagement,
	"Заведующий физиотерапевтическим отделением №2 - врач-физиотерапевт":                                        RoleMedicalManagement,
	"Заведующий хирургическим отделением №1 - врач-хирург":                                                      RoleMedicalManagement,
	"Заведующий хирургическим отделением №2 - врач-хирург":                                                      RoleMedicalManagement,
	"Заведующий хирургическим отделением №3 - врач-хирург":                                                      RoleMedicalManagement,
	"Заведующий центром амбулаторной онкологической помощи - врач-онколог":                                      RoleMedicalManagement,
	"Заведующий эндоскопическим отделением №1 - врач-эндоскопист":                                               RoleMedicalManagement,
	"Заведующий эндоскопическим отделением №2 - врач-эндоскопист":                                               RoleMedicalManagement,
	"Заведующий дневным стационаром - врач-терапевт":                                                            RoleMedicalManagement,
	"Заведующий рентгенологическим отделением №1 - врач-рентгенолог":                                            RoleMedicalManagement,
	"Заведующий рентгенологическим отделением №2 - врач-рентгенолог":                                            RoleMedicalManagement,
	"Заведующий эпидемиологическим отделом - врач-эпидемиолог":                                                  RoleMedicalManagement,

	// ВРАЧИ
	"врач-акушер-гинеколог":                     RoleDoctors,
	"врач-аллерголог-иммунолог":                 RoleDoctors,
	"врач-анестезиолог-реаниматолог":            RoleDoctors,
	"врач-бактериолог":                          RoleDoctors,
	"врач-гастроэнтеролог":                      RoleDoctors,
	"врач-дерматовенеролог":                     RoleDoctors,
	"врач-детский кардиолог":                    RoleDoctors,
	"врач-детский хирург":                       RoleDoctors,
	"врач-детский эндокринолог":                 RoleDoctors,
	"врач-диетолог":                             RoleDoctors,
	"врач-инфекционист":                         RoleDoctors,
	"врач-кардиолог":                            RoleDoctors,
	"врач-клинический фармаколог":               RoleDoctors,
	"врач-лаборант":                             RoleDoctors,
	"врач-методист":                             RoleDoctors,
	"врач-невролог":                             RoleDoctors,
	"врач-нейрохирург":                          RoleDoctors,
	"врач-неонатолог":                           RoleDoctors,
	"врач-нефролог":                             RoleDoctors,
	"врач-онколог":                              RoleDoctors,
	"врач-оториноларинголог":                    RoleDoctors,
	"врач-офтальмолог":                          RoleDoctors,
	"врач-педиатр":                              RoleDoctors,
	"врач-педиатр участковый":                   RoleDoctors,
	"врач-профпатолог":                          RoleDoctors,
	"врач-психиатр":                             RoleDoctors,
	"врач-психиатр-нарколог":                    RoleDoctors,
	"врач-пульмонолог":                          RoleDoctors,
	"врач-рентгенолог":                          RoleDoctors,
	"врач-рентгенолог (внешн. совм.)":           RoleDoctors,
	"врач-стажер":                               RoleDoctors,
	"врач-статистик":                            RoleDoctors,
	"врач-стоматолог-хирург":                    RoleDoctors,
	"врач-терапевт":                             RoleDoctors,
	"врач-терапевт участковый":                  RoleDoctors,
	"врач-токсиколог":                           RoleDoctors,
	"врач-травматолог-ортопед":                  RoleDoctors,
	"врач-уролог":                               RoleDoctors,
	"врач-физиотерапевт":                        RoleDoctors,
	"врач-хирург":                               RoleDoctors,
	"врач-эндокринолог":                         RoleDoctors,
	"врач-эндоскопист":                          RoleDoctors,
	"врач-эпидемиолог":                          RoleDoctors,
	"врач клинической лабораторной диагностики": RoleDoctors,
	"врач общей практики (семейный врач)":       RoleDoctors,
	"врач по лечебной физкультуре":              RoleDoctors,
	"врач по медицинской профилактике":          RoleDoctors,
	"врач по медицинской реабилитации":          RoleDoctors,
	"врач приемного отделения":                  RoleDoctors,
	"врач ультразвуковой диагностики":           RoleDoctors,
	"врач функциональной диагностики":           RoleDoctors,
	"биолог":  RoleDoctors,
	"логопед": RoleDoctors,
	"медицинский психолог": RoleDoctors,

	// СРЕДНИЙ МЕДИЦИНСКИЙ ПЕРСОНАЛ
	"акушерка":                                RoleNursingStaff,
	"фельдшер":                                RoleNursingStaff,
	"фельдшер-лаборант":                       RoleNursingStaff,
	"главная медицинская сестра":              RoleNursingStaff,
	"старшая медицинская сестра":              RoleNursingStaff,
	"старшая акушерка":                        RoleNursingStaff,
	"старшая операционная медицинская сестра": RoleNursingStaff,
	"медицинская сестра":                      RoleNursingStaff,
	"медицинская сестра перевязочной":         RoleNursingStaff,
	"медицинская сестра врача общей практики (семейного врача)": RoleNursingStaff,
	"медицинская сестра диетическая":                            RoleNursingStaff,
	"медицинская сестра палатная (постовая)":                    RoleNursingStaff,
	"медицинская сестра по массажу":                             RoleNursingStaff,
	"медицинская сестра по физиотерапии":                        RoleNursingStaff,
	"медицинская сестра приемного отделения":                    RoleNursingStaff,
	"медицинская сестра процедурной":                            RoleNursingStaff,
	"медицинская сестра стерилизационной":                       RoleNursingStaff,
	"медицинская сестра участковая":                             RoleNursingStaff,
	"медицинская сестра-анестезист":                             RoleNursingStaff,
	"медицинский брат палатный (постовой)":                      RoleNursingStaff,
	"медицинский брат приемного отделения":                      RoleNursingStaff,
	"операционная медицинская сестра":                           RoleNursingStaff,
	"инструктор по лечебной физкультуре":                        RoleNursingStaff,
	"инструктор-методист по лечебной физкультуре":               RoleNursingStaff,

	// МЛАДШИЙ МЕДИЦИНСКИЙ ПЕРСОНАЛ
	"младшая медицинская сестра по уходу за больными": RoleJuniorMedical,
	"санитарка":            RoleJuniorMedical,
	"санитарка (вн.совм.)": RoleJuniorMedical,
	"санитарка (рентген)":  RoleJuniorMedical,
	"сестра-хозяйка":       RoleJuniorMedical,

	// ТЕХНИЧЕСКИЙ ПЕРСОНАЛ
	"Главный инженер":    RoleTechnical,
	"инженер":            RoleTechnical,
	"инженер по ремонту": RoleTechnical,
	"инженер-энергетик":  RoleTechnical,
	"механик":            RoleTechnical,
	"слесарь по эксплуатации и ремонту газового оборудования":                 RoleTechnical,
	"слесарь по ремонту и обслуживанию систем вентиляции и кондиционирования": RoleTechnical,
	"слесарь-ремонтник": RoleTechnical,
	"слесарь-сантехник": RoleTechnical,
	"электромеханик по ремонту и обслуживанию медицинского оборудования":         RoleTechnical,
	"старший электромеханик по ремонту и обслуживанию медицинского оборудования": RoleTechnical,
	"электромонтер по ремонту и обслуживанию электрооборудования":                RoleTechnical,
	"электромонтер линейных сооружений телефонной связи и радиофикации":          RoleTechnical,
	"электрогазосварщик":          RoleTechnical,
	"каменщик":                    RoleTechnical,
	"маляр":                       RoleTechnical,
	"плотник":                     RoleTechnical,
	"штукатур":                    RoleTechnical,
	"облицовщик-плиточник":        RoleTechnical,
	"заточник":                    RoleTechnical,
	"Ведущий инженер":             RoleTechnical,
	"Инженер-сметчик I категории": RoleTechnical,
	"Мастер службы":               RoleTechnical,
	"Мастер участка":              RoleTechnical,
	"Оператор агрегата обработки отходов": RoleTechnical,

	// ЭКОНОМИКА И БУХГАЛТЕРИЯ
	"главный бухгалтер":                     RoleEconomic,
	"заместитель главного бухгалтера":       RoleEconomic,
	"ведущий бухгалтер":                     RoleEconomic,
	"ведущий бухгалтер-руководитель группы": RoleEconomic,
	"бухгалтер I категории":                 RoleEconomic,
	"ведущий экономист":                     RoleEconomic,
	"экономист":                             RoleEconomic,
	"экономист I категории":                 RoleEconomic,
	"экономист II категории":                RoleEconomic,
	"кассир":               RoleEconomic,
	"старший кассир":       RoleEconomic,
	"Ведущий юрисконсульт": RoleEconomic,

	// IT-СПЕЦИАЛИСТЫ
	"ведущий инженер-программист": RoleIT,
	"инженер-программист":         RoleIT,
	"техник-программист":          RoleIT,
	"оператор электронно-вычислительных и вычислительных машин": RoleIT,
	"оператор копировальных и множительных машин":               RoleIT,
	"специалист по защите информации":                           RoleIT,

	// КАДРЫ И АДМИНИСТРАТИВНЫЙ ПЕРСОНАЛ
	"начальник отдела кадров":                        RoleHR,
	"специалист по кадрам":                           RoleHR,
	"инспектор по контролю за исполнением поручений": RoleHR,
	"делопроизводитель":                              RoleHR,
	"архивариус":                                     RoleHR,
	"администратор":                                  RoleHR,
	"старший администратор":                          RoleHR,
	"секретарь":                                      RoleHR,
	"Специалист по связям с общественностью":         RoleHR,

	// ОБСЛУЖИВАЮЩИЙ ПЕРСОНАЛ
	"заведующий прачечной":                             RoleSupport,
	"заведующий производством":                         RoleSupport,
	"заведующий складом":                               RoleSupport,
	"заведующий хозяйством":                            RoleSupport,
	"начальник по административно-хозяйственной части": RoleSupport,
	"начальник хозяйственной службы":                   RoleSupport,
	"повар":                       RoleSupport,
	"пекарь":                      RoleSupport,
	"буфетчик":                    RoleSupport,
	"кухонный рабочий":            RoleSupport,
	"мойщик посуды":               RoleSupport,
	"мойщик посуды и ампул":       RoleSupport,
	"уборщик служебных помещений": RoleSupport,
	"уборщик территорий":          RoleSupport,
	"гардеробщик":                 RoleSupport,
	"лифтер":                      RoleSupport,
	"грузчик":                     RoleSupport,
	"подсобный рабочий":           RoleSupport,
	"кладовщик":                   RoleSupport,
	"курьер":                      RoleSupport,
	"старший курьер":              RoleSupport,
	"диспетчер":                   RoleSupport,
	"швея":                        RoleSupport,
	"фасовщица":                   RoleSupport,
	"упаковщик (вручную)":         RoleSupport,
	"оператор стиральных машин":   RoleSupport,
	"Агент по снабжению":          RoleSupport,
	"Ведущий специалист":          RoleSupport,
	"Водитель автомобиля":         RoleSupport,
	"Дезинфектор":                 RoleSupport,
	"Кастелянша":                  RoleSupport,
	"Контролер контрольно-пропускного пункта": RoleSupport,
	"Лаборант": RoleSupport,
	"Медицинский дезинфектор":                RoleSupport,
	"Медицинский статистик":                  RoleSupport,
	"Методист":                               RoleSupport,
	"Начальник гаража":                       RoleSupport,
	"Начальник контрольно-пропускной службы": RoleSupport,
	"Заместитель начальника контрольно-пропускной службы": RoleSupport,
	"Преподаватель":                           RoleSupport,
	"Провизор":                                RoleSupport,
	"Провизор-аналитик":                       RoleSupport,
	"Провизор-технолог":                       RoleSupport,
	"Рентгенолаборант":                        RoleSupport,
	"Рентгенолаборант (внешн. совм.)":         RoleSupport,
	"Специалист гражданской обороны":          RoleSupport,
	"Специалист по закупкам":                  RoleSupport,
	"Специалист по охране труда":              RoleSupport,
	"Специалист по пожарной безопасности":     RoleSupport,
	"Специалист по социальной работе":         RoleSupport,
	"Статистик":                               RoleSupport,
	"Стерилизаторщик материалов и препаратов": RoleSupport,
	"Техник":    RoleSupport,
	"Фармацевт": RoleSupport,
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

		role := GetRole(employees[i].Dol_Name, db)
		if role.ID == 0 {
			continue
		}
		AddRole(role.ID, employees[i].UserID, db)

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

func GetRole(position string, db *sqlx.DB) Role {
	r, exists := RoleMapping[position]
	if !exists {
		return Role{}
	}

	var role Role

	query := `
		SELECT id, name from roles where name=$1
	`

	err := db.Get(&role, query, r)
	if err != nil {
		log.Printf("Ошибка при получении роли: %v", err)
		return Role{}
	}

	return role
}

func AddRole(roleId, userId int, db *sqlx.DB) {
	query := `
		INSERT INTO user_role (user_id, role_id) VALUES ($1, $2)
	`

	_, err := db.Exec(query, userId, roleId)
	if err != nil {
		log.Printf("Ошибка при добавлении роли: %v", err)
	}
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
