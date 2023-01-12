package services

import (
	"fmt"
	"library-api/dto"
	"library-api/httperror"
	"library-api/models"
	repository "library-api/repositories"
	"time"
)

type BorrowingService interface {
	BorrowBook(borrowingReq *dto.BorrowingReq) (*models.BorrowingRecord, error)
	ReturnBook(id int) (*models.BorrowingRecord, error)
}

type borrowingService struct {
	borrowingRepository repository.BorrowingRepository
	bookRepository      repository.BookRepository
}

type BRWOpts struct {
	BorrowingRepository repository.BorrowingRepository
	BookRepository      repository.BookRepository
}

func NewBorrowingService(opts *BRWOpts) BorrowingService {
	return &borrowingService{
		borrowingRepository: opts.BorrowingRepository,
		bookRepository:      opts.BookRepository,
	}
}

func (b *borrowingService) BorrowBook(borrowingReq *dto.BorrowingReq) (*models.BorrowingRecord, error) {
	fmt.Println("===", b.bookRepository)
	book, err := b.bookRepository.FindBook(borrowingReq.BookID)
	if err != nil {
		return nil, httperror.BadRequestError("book not found", "")
	}
	if book.Stock <= 0 {
		return nil, httperror.BadRequestError("book out of stock", "")
	}
	rec := models.BorrowingRecord{
		UserID:        borrowingReq.UserID,
		BookID:        borrowingReq.BookID,
		BorrowingDate: time.Now(),
		ReturningDate: nil,
	}
	record, err := b.borrowingRepository.Create(rec)
	if err != nil {
		return nil, httperror.InternalServerError("failed to create record")
	}
	err = b.bookRepository.DecreaseStock(book.ID)
	return record, err
}

func (b *borrowingService) ReturnBook(recordID int) (*models.BorrowingRecord, error) {
	record, err := b.borrowingRepository.ReturnBook(recordID)
	if err != nil {
		return nil, httperror.InternalServerError("failed to return book")
	}
	err = b.bookRepository.IncreaseStock(record.BookID)
	return record, err
}
