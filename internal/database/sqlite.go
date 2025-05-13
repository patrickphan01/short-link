package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func NewSQLite(){
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	
	if err != nil {
		panic("failed to connet to the database")
	}

}