package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Store *gorm.DB

func InitPostgres(dbname string) error {
	db, err := gorm.Open("postgres", dbname)
	if err != nil {
		return err
	}
	Store = db
	return nil
}
