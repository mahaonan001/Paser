package model

import (
	"gorm.io/gorm"
)

type PaperNew struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Titles    string `gorm:"type:varchar(20);not null"`
	Questions string `gorm:"type:TEXT(65535);not null"`
}
