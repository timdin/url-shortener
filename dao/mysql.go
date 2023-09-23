package dao

import (
	"url-shortener/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDao struct {
	db *gorm.DB
}

func InitDB(dbConfig string) (*MysqlDao, error) {
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	migrateErr := db.AutoMigrate(model.URL{})
	if migrateErr != nil {
		return nil, err
	}
	return &MysqlDao{
		db: db,
	}, nil
}

// TODO: make it upsert with shortURL as primary key
func (dao *MysqlDao) WriteURLRecord(entity *model.URL) error {
	return dao.db.Create(entity).Error
}
