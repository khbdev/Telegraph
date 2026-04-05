package models

import "gorm.io/gorm"


type URL struct {
	gorm.Model
	     string `gorm:"type:varchar(255);not null"`
}