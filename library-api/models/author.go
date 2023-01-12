package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Name       string `json:"name"`
}
