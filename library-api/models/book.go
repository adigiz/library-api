package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model  `json:"-"`
	ID          uint   `gorm:"primarykey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	AuthorID    uint   `json:"authorID"`
	Author      Author `json:"author"`
	Stock       int    `json:"stock"`
}
