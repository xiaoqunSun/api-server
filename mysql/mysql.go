package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(dataSourceName string) error {
	var err error

	if db == nil {
		db, err = sql.Open("mysql", dataSourceName)
	}
	return err

}

func DB() *sql.DB {
	return db
}
