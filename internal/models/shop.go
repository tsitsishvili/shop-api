package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name   string `json:"name"`
	APIKey string `gorm:"unique"`
}
