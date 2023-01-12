package models

import (
	"gorm.io/gorm"
	"time"
)

type BorrowingRecord struct {
	gorm.Model    `json:"-"`
	ID            uint       `json:"id"`
	UserID        int        `json:"user_id"`
	BookID        int        `json:"book_id"`
	Status        string     `json:"status"`
	BorrowingDate time.Time  `json:"borrowing_date"`
	ReturningDate *time.Time `json:"returning_date"`
}
