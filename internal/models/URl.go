package models

import "gorm.io/gorm"


type URL struct {
	gorm.Model
	URL    string `gorm:"type:varchar(255);not null"`

	DataID uint  
	Data   Data   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}