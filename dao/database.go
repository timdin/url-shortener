package dao

import (
	"time"
	"url-shortener/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBinterface interface {
	WriteDB(entity *model.URL) error
	QueryDB(shortURL string) (*model.URL, error)
}

type DBImpl struct {
	db *gorm.DB
}

func initDB(dbConfig string) (*DBImpl, error) {
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	migrateErr := db.AutoMigrate(model.URL{})
	if migrateErr != nil {
		return nil, err
	}
	return &DBImpl{
		db: db,
	}, nil
}

func (db *DBImpl) WriteDB(entity *model.URL) error {
	// support upsert with same url to update the expiration
	err := db.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "short_url"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"expires_at": entity.ExpiresAt,
		}),
	}).Create(entity).Error
	return err
}

func (db *DBImpl) QueryDB(shortURL string) (*model.URL, error) {
	res := &model.URL{}
	if err := db.db.Where("short_url = ?", shortURL).Where("expires_at > ? OR expires_at IS NULL ", time.Now()).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
