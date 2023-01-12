package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"-"`
}
