package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShortLink struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Origin    string    `gorm:"nottype:text;not null"`
	Path      string    `gorm:"type:varchar(255)"`
	ShortUrl  string    `gorm:"type:text;not null"`
	RequestIP string    `gorm:"type:varchar(45)"`
	// ExpireTime time.Time `gorm:"not null"`
}

func (s *ShortLink) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
