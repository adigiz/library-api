package services

import (
	"git.garena.com/sea-labs-id/trainers/library-api/dto"
	"git.garena.com/sea-labs-id/trainers/library-api/httperror"
	"git.garena.com/sea-labs-id/trainers/library-api/models"
	repository "git.garena.com/sea-labs-id/trainers/library-api/repositories"
)

type BookService interface {
	GetBooks() ([]*models.Book, error)
	AddBook(payload *dto.BookReq) (*dto.BookRes, error)
}

type BSOpts struct {
	BookRepository repository.BookRepository
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(o *BSOpts) BookService {
	return &bookService{
		bookRepository: o.BookRepository,
	}
}

func (b *bookService) GetBooks() ([]*models.Book, error) {
	return b.bookRepository.FindBooks()
}

func (b *bookService) AddBook(payload *dto.BookReq) (*dto.BookRes, error) {
	book := models.Book{
		Title:       payload.Title,
		Description: payload.Description,
		Cover:       payload.Cover,
		AuthorID:    payload.AuthorID,
	}
	insertedBook, rowsAffected, err := b.bookRepository.Save(&book)
	if err == nil && rowsAffected == 0 {
		return new(dto.BookRes), httperror.BadRequestError("Duplicate book", "DUPLICATE_BOOK")
	}
	return new(dto.BookRes).FromBook(insertedBook), err
}
