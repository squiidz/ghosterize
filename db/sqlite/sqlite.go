package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Store *gorm.DB

func InitSqlite(dbname string) error {
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		return err
	}
	Store = db
	return nil
}
