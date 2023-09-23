package model

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	ShortURL  string     `gorm:"unique" json:"short_url"`
	LongURL   string     `json:"long_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
}
