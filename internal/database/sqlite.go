package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/patrickphan01/short-link/internal/models"
)


func NewSQLite(){
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	
	if err != nil {
		panic("failed to connet to the database")
	}
	db.AutoMigrate(&models.ShortLink{})
}