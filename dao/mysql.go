package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dbConfig string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	return db, err
}
