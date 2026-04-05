package models

import "gorm.io/gorm"


type URL struct {
	gorm.Model
	Title     string `gorm:"type:varchar(255);not null"`
	
}