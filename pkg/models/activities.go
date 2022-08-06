package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name        string  `gorm:"uniqueIndex" json:"name"`
	Description string  `json:"description"`
	Effort      float64 `json:"effort"`
}
