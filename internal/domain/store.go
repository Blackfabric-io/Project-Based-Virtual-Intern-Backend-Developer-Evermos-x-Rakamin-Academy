package domain

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID uint
	Name   string
	// Field lain sesuai kebutuhan
}
