package dao

import (
	"log"
	"time"
	"url-shortener/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlDao struct {
	db *gorm.DB
}

func InitDB(dbConfig string) (*MysqlDao, error) {
	log.Println(dbConfig)
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

func (dao *MysqlDao) WriteURLRecord(entity *model.URL) error {
	// support upsert with same url to update the expiration
	err := dao.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "short_url"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"expires_at": entity.ExpiresAt,
		}),
	}).Create(entity).Error
	return err
}

func (dao *MysqlDao) QueryURLRecord(shortURL string) (*model.URL, error) {
	res := &model.URL{}
	// TODO: add cache here
	if err := dao.db.Where("short_url = ?", shortURL).Where("expires_at > ? OR expires_at IS NULL ", time.Now()).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
