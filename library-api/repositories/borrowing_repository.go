package repository

import (
	"library-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type BorrowingRepository interface {
	Create(b models.BorrowingRecord) (*models.BorrowingRecord, error)
	ReturnBook(recordID int) (*models.BorrowingRecord, error)
}

type borrowingRepository struct {
	db *gorm.DB
}

type BRWOpts struct {
	DB *gorm.DB
}

func NewBorrowingRepository(opts *BRWOpts) BorrowingRepository {
	return &borrowingRepository{db: opts.DB}
}

func (u *borrowingRepository) Create(b models.BorrowingRecord) (*models.BorrowingRecord, error) {
	b.Status = "BORROWED"
	result := u.db.Create(&b)
	return &b, result.Error
}

func (u *borrowingRepository) ReturnBook(recordID int) (*models.BorrowingRecord, error) {
	record := models.BorrowingRecord{ID: uint(recordID)}
	now := time.Now()
	updateData := models.BorrowingRecord{ReturningDate: &now, Status: "RETURNED"}
	result := u.db.Model(&record).Clauses(clause.Returning{}).Updates(updateData)
	return &record, result.Error
}
