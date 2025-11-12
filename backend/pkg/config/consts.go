package config

import (
	"fmt"
	"portal/internal/entity"

	"github.com/jmoiron/sqlx"
)

var InitialRoles = []entity.RoleCreate{
	{Name: "Администрация"},
	{Name: "Медицинские руководители"},
	{Name: "Врачи"},
	{Name: "Средний медицинский персонал"},
	{Name: "Младший медицинский персонал"},
	{Name: "Технический персонал"},
	{Name: "Экономика и бухгалтерия"},
	{Name: "IT-специалисты"},
	{Name: "Кадры и административный персонал"},
	{Name: "Обслуживающий персонал"},

	// Роли для массовых операций
	{Name: "Массовые операции - Трудоустройство"},
	{Name: "Массовые операции - Увольнение"},

	// Роли для технической поддержки
	{Name: "Техподдержка - Компьютеры"},
	{Name: "Техподдержка - Принтеры"},
	{Name: "Техподдержка - Программы"},
	{Name: "Техподдержка - Порталы"},
	{Name: "Техподдержка - Ариадна"},
}

var InitialTasks = []entity.TaskCreate{
	{Name: "Трудоустройство", Type: "mass"},
	{Name: "Увольнение", Type: "mass"},
	{Name: "Не работает компьютер", Type: "support"},
	{Name: "Не работает принтер", Type: "support"},
	{Name: "Закончился тонер", Type: "support"},
	{Name: "Не работает программа", Type: "support"},
	{Name: "Не работает портал", Type: "support"},
	{Name: "Не работает Ариадна", Type: "support"},
}

var InitialBranches = []entity.BranchCreate{
	{Name: "Островского"},
	{Name: "Александрова"},
	// ... все филиалы
}

var InitialPrinters = []entity.PrinterCreate{
	{Value: "brother-dcp-l2520wr", Name: "Brother DCP-L2520wr"},
	{Value: "brother-hl-1110r", Name: "brother hl-1110r"},
	{Value: "brother-hl-2132r", Name: "Brother HL 2132R"},
	{Value: "brother-mfc-l2700dnr", Name: "Brother MFC L2700DNR"},
	{Value: "canon-i-sensys-lbp6030b", Name: "Canon i-SENSYS LBP6030B"},
	{Value: "canon-i-sensys-mf4018", Name: "Canon i-SENSYS MF4018"},
	{Value: "canon-l11121e-lbp-2900", Name: "Canon L11121E /LBP-2900"},
	{Value: "canon-laserbase-mf3010", Name: "Canon LaserBase MF3010"},
	{Value: "canon-laserbase-mf4410-i-sensys", Name: "Canon LaserBase MF4410 i-Sensys"},
	{Value: "canon-lbp-2900", Name: "Canon LBP-2900"},
	{Value: "canon-lbp-3010", Name: "Canon LBP 3010"},
	{Value: "canon-lbp-6030", Name: "Canon LBP 6030"},
	{Value: "canon-mf-735cx-i-sensys-snyat", Name: "Canon MF 735Cx i-Sensys/снят"},
	{Value: "copir-xerox-copycentre-c118-a3", Name: "Копир Xerox CopyCentre C118 A3"},
	{Value: "epson-l-3101", Name: "Epson L 3101"},
	{Value: "epson-l-4160", Name: "Epson L 4160"},
	{Value: "hp-1018", Name: "HP 1018"},
	{Value: "hp-color-laserjet-179fnw-mfp", Name: "HP Color LaserJet 179fnw MFP"},
	{Value: "hp-color-laserjet-cm1415fn-pro", Name: "HP Color LaserJet CM1415fn Pro"},
	{Value: "hp-laser-135a-mfp", Name: "HP Laser 135a MFP"},
	{Value: "hp-laserjet-1010", Name: "HP LaserJet 1010"},
	{Value: "hp-laserjet-1022", Name: "HP LaserJet 1022"},
	{Value: "hp-laserjet-m1132-mfp", Name: "HP LaserJet M1132 MFP"},
	{Value: "hp-laserjet-m1212nf-mfp-pro", Name: "HP LaserJet M1212nf MFP Pro"},
	{Value: "hp-laserjet-m28-mfp-pro", Name: "HP LaserJet M28 MFP Pro"},
	{Value: "hp-laserjet-m404dn-pro", Name: "HP LaserJet M404dn Pro"},
	{Value: "hp-laserjet-p103-107", Name: "HP LaserJet P103-107"},
	{Value: "hp-laserjet-p1102", Name: "HP LaserJet P1102"},
	{Value: "hp-laserjet-p1505", Name: "HP LaserJet P1505"},
	{Value: "hp-laserjet-p2050", Name: "HP LaserJet P2050"},
	{Value: "hp-laserjet-pro-m125ra", Name: "HP LaserJetPro M125ra"},
	{Value: "hp-laserjet-pro-m203dn", Name: "HP LaserJet Pro M203dn"},
	{Value: "hp-laserjet-pro-m402dne", Name: "HP LaserJet Pro M402dne"},
	{Value: "hp-laserjetm207-m212", Name: "hp laserjet m207-m212"},
	{Value: "hp-laserjetm211dw", Name: "hp laserjet m211dw"},
	{Value: "konica-minolta-bizhub-185", Name: "Konica Minolta Bizhub 185"},
	{Value: "kyocera-ecosys-m4125idn", Name: "Kyosera ECOSYS M4125idn"},
	{Value: "kyocera-fs-1035", Name: "Kyosera FS 1035"},
	{Value: "kyocera-fs-1040", Name: "Kyosera FS 1040"},
	{Value: "kyocera-fs-1100", Name: "Kyosera FS 1100"},
	{Value: "kyocera-fs-1110", Name: "Kyosera FS 1110"},
	{Value: "kyocera-fs-1120d", Name: "Kyosera FS 1120D"},
	{Value: "kyocera-fs-1128", Name: "Kyosera FS 1128"},
	{Value: "kyocera-fs-1300d", Name: "Kyosera FS 1300D"},
	{Value: "kyocera-km-2050-a3", Name: "Kyosera KM 2050 (A3)"},
	{Value: "kyocera-m2040dn", Name: "Kyosera M2040dn"},
	{Value: "kyocera-m2235dn", Name: "Kyosera M2235dn"},
	{Value: "kyocera-m2540dn", Name: "Kyosera M2540dn"},
	{Value: "kyocera-p2035d", Name: "Kyosera P2035d"},
	{Value: "kyocera-p2040dn", Name: "Kyosera P2040dn"},
	{Value: "kyocera-p2235dn", Name: "Kyosera P2235dn"},
	{Value: "kyocera-p2335d", Name: "Kyosera P2335d"},
	{Value: "kyocera-p2335dn", Name: "Kyosera P2335dn"},
	{Value: "kyocera-taskalfa-1800-a3", Name: "Kyosera TASKALFA 1800 (A3)"},
	{Value: "pantum-bm5100adn", Name: "Pantum BM5100ADN"},
	{Value: "pantum-m6502w", Name: "Pantum M6502W"},
	{Value: "pantum-m6507", Name: "Pantum M6507"},
	{Value: "pantum-p2200", Name: "Pantum P2200"},
	{Value: "pantum-p2500nw", Name: "Pantum P2500NW"},
	{Value: "rizograf-ricoh-jp750", Name: "Ризограф Ricoh JP750"},
	{Value: "ricoh-priport-jp750", Name: "Ricoh Priport JP750"},
	{Value: "samsung-ml-1860", Name: "Samsung ML 1860"},
	{Value: "samsung-ml-2160", Name: "Samsung ML 2160"},
	{Value: "samsung-ml1665", Name: "Samsung ML1665"},
	{Value: "samsung-ml2160", Name: "Samsung ML2160"},
	{Value: "samsung-ml3750nd", Name: "Samsung ML3750ND"},
	{Value: "samsung-scx-4220", Name: "Samsung SCX 4220"},
	{Value: "samsung-xpress-m2070", Name: "Samsung Xpress M2070"},
	{Value: "xerox-3020", Name: "XEROX 3020"},
	{Value: "xerox-3025", Name: "Xerox 3025"},
	{Value: "xerox-b210", Name: "XEROX B210"},
	{Value: "xerox-phaser-3117", Name: "Xerox Phaser 3117"},
}

var InitialRooms = []entity.RoomCreate{
	{Value: "kab", Name: "Кабинет"},
	{Value: "post", Name: "Пост"},
	{Value: "ord", Name: "Ординаторская"},
	{Value: "reg", Name: "Регистратура"},
	{Value: "priem", Name: "Приемное отделение"},
	{Value: "palata", Name: "Палата"},
	{Value: "hall", Name: "Коридор"},
}

func SeedInitialData(db *sqlx.DB) error {
	if err := seedBranches(db); err != nil {
		return fmt.Errorf("seeding departments: %w", err)
	}

	if err := seedRoles(db); err != nil {
		return fmt.Errorf("seeding roles: %w", err)
	}

	if err := seedTasks(db); err != nil {
		return fmt.Errorf("seeding tasks: %w", err)
	}

	if err := seedPrinters(db); err != nil {
		return fmt.Errorf("seeding printers: %w", err)
	}

	if err := seedRooms(db); err != nil {
		return fmt.Errorf("seeding rooms: %w", err)
	}

	return nil
}

func seedBranches(db *sqlx.DB) error {
	for _, dept := range InitialBranches {
		var existing entity.Branch
		err := db.Get(&existing, "SELECT * FROM branches WHERE name = $1", dept.Name)
		if err != nil {
			// Отделение не существует, добавляем
			_, err := db.NamedExec(`
                INSERT INTO branches (name) 
                VALUES (:name)
            `, dept)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func seedRoles(db *sqlx.DB) error {
	for _, role := range InitialRoles {
		var existing entity.Role
		err := db.Get(&existing, "SELECT * FROM roles WHERE name = $1", role.Name)
		if err != nil {
			// Роль не существует, добавляем
			_, err := db.NamedExec(`
                INSERT INTO roles (name) 
                VALUES (:name)
            `, role)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func seedTasks(db *sqlx.DB) error {
	for _, task := range InitialTasks {
		var existing entity.Task
		err := db.Get(&existing, "SELECT * FROM task_list WHERE name = $1", task.Name)
		if err != nil {
			// Задача не существует, добавляем
			_, err := db.NamedExec(`
                INSERT INTO task_list (name, type) 
                VALUES (:name, :type)
            `, task)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func seedPrinters(db *sqlx.DB) error {
	for _, printer := range InitialPrinters {
		var existing entity.Printer
		err := db.Get(&existing, "SELECT * FROM printers WHERE value = $1", printer.Value)
		if err != nil {
			// Принтер не существует, добавляем
			_, err := db.NamedExec(`
                INSERT INTO printers (value, name) 
                VALUES (:value, :name)
            `, printer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func seedRooms(db *sqlx.DB) error {
	for _, room := range InitialRooms {
		var existing entity.Room
		err := db.Get(&existing, "SELECT * FROM rooms WHERE value = $1", room.Value)
		if err != nil {
			// Помещение не существует, добавляем
			_, err := db.NamedExec(`
                INSERT INTO rooms (value, name) 
                VALUES (:value, :name)
            `, room)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
