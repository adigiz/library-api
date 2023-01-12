package services_test

import (
	"errors"
	"git.garena.com/sea-labs-id/trainers/library-api/httperror"
	"testing"

	"git.garena.com/sea-labs-id/trainers/library-api/dto"
	"git.garena.com/sea-labs-id/trainers/library-api/mocks"
	"git.garena.com/sea-labs-id/trainers/library-api/models"
	. "git.garena.com/sea-labs-id/trainers/library-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	book = &models.Book{
		ID:          1,
		Title:       "Clean Code",
		Description: "Clean code for developer",
		Cover:       "https://hehe.com",
		AuthorID:    1,
	}
	bookRes = &dto.BookRes{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Cover:       book.Cover,
		AuthorID:    book.AuthorID,
	}
	books = []*models.Book{book}
)

func TestBook_GetBooks(t *testing.T) {
	t.Run("should return list of book when no error", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)
		opts := &BSOpts{BookRepository: mockRepo}
		s := NewBookService(opts)
		mockRepo.On("FindBooks").Return(books, nil)

		actualBooks, err := s.GetBooks()

		assert.Equal(t, books, actualBooks)
		assert.Nil(t, err)
	})

	t.Run("should return empty list and error when error get books from repository", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)
		opts := &BSOpts{BookRepository: mockRepo}
		s := NewBookService(opts)
		dbError := errors.New("unable to connect to db")
		var expectedBooks []*models.Book
		mockRepo.On("FindBooks").Return(expectedBooks, dbError)

		actualBooks, err := s.GetBooks()

		assert.Equal(t, dbError, err)
		assert.Equal(t, expectedBooks, actualBooks)
	})
}

func TestBook_AddBook(t *testing.T) {
	t.Run("should return inserted book when no error occurred", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)
		payload := dto.BookReq{
			Title:       "Clean Code",
			Description: "Clean code for developer",
			Cover:       "https://hehe.com",
		}
		mockRepo.On("Save", mock.AnythingOfTypeArgument("*models.Book")).Return(book, 1, nil)
		opts := &BSOpts{BookRepository: mockRepo}
		s := NewBookService(opts)

		actualBooks, err := s.AddBook(&payload)

		assert.Equal(t, bookRes, actualBooks)
		assert.Nil(t, err)
	})

	t.Run("should return error DUPLICATE_BOOK when rowsAffected is 0", func(t *testing.T) {
		mockRepo := new(mocks.BookRepository)
		payload := dto.BookReq{
			Title:       "Clean Code",
			Description: "Clean code for developer",
			Cover:       "https://hehe.com",
		}
		rowsAffected := 0
		mockRepo.On("Save", mock.AnythingOfTypeArgument("*models.Book")).Return(book, rowsAffected, nil)
		expectedErr := httperror.BadRequestError("Duplicate book", "DUPLICATE_BOOK")
		opts := &BSOpts{BookRepository: mockRepo}
		s := NewBookService(opts)

		_, err := s.AddBook(&payload)

		assert.Equal(t, expectedErr, err)
	})
}
