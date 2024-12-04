package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Phone    string `json:"phone" gorm:"unique"`
	isAdmin  string `json:"is_admin"`
}
