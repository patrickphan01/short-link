package models

import "github.com/google/uuid"

type Users struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string `gorm:"type:varchar(255);not null;unique"`
	IsActive bool `gorm:"type:boolean;default=false"`
	Password string `gorm:"type:varchar(255);not null"`
	CountUrl int `gorm:"type:int;default:0"`
	// Package  int 
	
}