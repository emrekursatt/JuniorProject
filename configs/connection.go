package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(config Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.Dbname)

	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		println(err.Error())
		return nil, err
	}

	if err = db.Ping(); err != nil {
		println(err)
		return nil, err
	}

	return db, nil
}

var cfg = DefaultConfig()
var DB, err = NewConnection(cfg)

func CreateTaskTable(db *sql.DB, tableName string) (*sql.DB, error) {
	isTableExist, _ := isAlreadyTableExist(db, tableName)
	if isTableExist {
		crateTable := fmt.Sprintf("CREATE TABLE `%s`.`%s` "+
			"( `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,  "+
			"`code` VARCHAR(15) NOT NULL,\n  "+
			"`title` VARCHAR(45) NOT NULL,\n  "+
			"`description` VARCHAR(45) NULL,\n  "+
			"`status` VARCHAR(45) NOT NULL,\n  "+
			"PRIMARY KEY (`id`),\n  "+
			"UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,  "+
			"UNIQUE INDEX `title_UNIQUE` "+
			"(`title` ASC) VISIBLE) "+
			"ENGINE = InnoDB ", cfg.Dbname, tableName)

		_, err = db.Exec(crateTable)
	}
	return db, nil
}

func CreateUserTable(db *sql.DB, tableName string) (*sql.DB, error) {
	isTableExist, _ := isAlreadyTableExist(db, tableName)
	if isTableExist {
		crateTable := fmt.Sprintf("CREATE TABLE `%s`.`%s` "+
			"(  `id` INT(11) NOT NULL AUTO_INCREMENT,  "+
			"`username` VARCHAR(45) NOT NULL,  "+
			"`password` VARCHAR(45) NOT NULL,  "+
			"`email` VARCHAR(100) NOT NULL,  "+
			"`phone_number` BIGINT(11) UNSIGNED ZEROFILL NOT NULL, "+
			" PRIMARY KEY (`id`, `username`), UNIQUE INDEX `id_UNIQUE` "+
			"(`id` ASC) VISIBLE, UNIQUE INDEX `username_UNIQUE` "+
			"(`phone_number` ASC) VISIBLE)", cfg.Dbname, tableName)

		_, err = db.Exec(crateTable)
	}
	return db, nil
}

func isAlreadyTableExist(db *sql.DB, tableName string) (bool, error) {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*)
              FROM information_schema.tables
              WHERE table_name = '%s'`, tableName)

	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	return true, nil
}
