package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name   string
	Rating int16
}
