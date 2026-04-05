package models

import "gorm.io/gorm"


type URL struct {
	gorm.Model
	URL string `gorm:"type:varchar(255);not null"`
}