package handler_test

import (
	"encoding/json"
	"errors"
	"library-api/dto"
	"library-api/httperror"
	"library-api/mocks"
	"library-api/models"
	"library-api/server"
	"library-api/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

var (
	bookPayload = dto.BookReq{
		Title:       "Clean Code",
		Description: "Clean code for developer",
		Cover:       "https://hehe.com",
		AuthorID:    1,
	}
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

func TestHandler_GetBooks(t *testing.T) {
	t.Run("should return code 200 with list of books when no error", func(t *testing.T) {
		mockBService := new(mocks.BookService)
		mockBService.On("GetBooks").Return(books, nil)
		opts := &server.RouterOpts{
			BookService: mockBService,
		}
		expectedRes, _ := json.Marshal(books)

		req, _ := http.NewRequest("GET", "/books", nil)
		_, rec := testutils.ServeReq(opts, req)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedRes), rec.Body.String())
	})

	t.Run("should return code 500 when error", func(t *testing.T) {
		mockBService := new(mocks.BookService)
		mockError := errors.New("failed to connect to DB")
		mockBService.On("GetBooks").Return(nil, mockError)
		opts := &server.RouterOpts{
			BookService: mockBService,
		}

		req, _ := http.NewRequest("GET", "/books", nil)
		_, rec := testutils.ServeReq(opts, req)

		assert.Equal(t, 500, rec.Code)
	})
}

func TestHandler_PostBook(t *testing.T) {
	t.Run("should return code 201 when book inserted", func(t *testing.T) {
		mockBService := new(mocks.BookService)
		mockBService.On("AddBook", mock.AnythingOfTypeArgument("*dto.BookReq")).Return(bookRes, nil)
		opts := &server.RouterOpts{
			BookService: mockBService,
		}
		body := testutils.MakeRequestBody(bookPayload)
		expectedRes, _ := json.Marshal(bookRes)

		req, _ := http.NewRequest("POST", "/books", body)
		_, rec := testutils.ServeReq(opts, req)

		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, string(expectedRes), rec.Body.String())
	})

	t.Run("should return code 400 BAD_REQUEST when required body is not exist", func(t *testing.T) {
		mockBService := new(mocks.BookService)
		opts := &server.RouterOpts{
			BookService: mockBService,
		}
		body := testutils.MakeRequestBody(dto.BookReq{
			Cover: "https://hehe.com",
		})
		expectedErrBody := httperror.AppError{
			StatusCode: 400,
			Code:       "BAD_REQUEST",
		}

		req, _ := http.NewRequest("POST", "/books", body)
		_, rec := testutils.ServeReq(opts, req)
		var actualResBody httperror.AppError
		_ = json.Unmarshal(rec.Body.Bytes(), &actualResBody)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, expectedErrBody.Code, actualResBody.Code)
		assert.Equal(t, expectedErrBody.StatusCode, actualResBody.StatusCode)
		assert.NotEmpty(t, actualResBody.Message)
	})

	t.Run("should forward http error when service return app error", func(t *testing.T) {
		mockBService := new(mocks.BookService)
		expectedErrBody := httperror.AppError{
			StatusCode: 400,
			Code:       "DUPLICATE_BOOK",
		}
		appError := httperror.BadRequestError("Duplicate book", expectedErrBody.Code)
		mockBService.On("AddBook", mock.AnythingOfTypeArgument("*dto.BookReq")).Return(nil, appError)
		opts := &server.RouterOpts{
			BookService: mockBService,
		}
		body := testutils.MakeRequestBody(bookPayload)

		req, _ := http.NewRequest("POST", "/books", body)
		_, rec := testutils.ServeReq(opts, req)
		var actualResBody httperror.AppError
		_ = json.Unmarshal(rec.Body.Bytes(), &actualResBody)

		assert.Equal(t, 400, rec.Code)
		assert.Equal(t, expectedErrBody.Code, actualResBody.Code)
		assert.Equal(t, expectedErrBody.StatusCode, actualResBody.StatusCode)
		assert.NotEmpty(t, actualResBody.Message)
	})

	t.Run("should return code 500 when unknown error occurred", func(t *testing.T) {
		mockBService := new(mocks.BookService)
		expectedErrBody := httperror.AppError{
			StatusCode: 500,
			Code:       "INTERNAL_SERVER_ERROR",
		}
		err := errors.New("internal server err")
		mockBService.On("AddBook", mock.AnythingOfTypeArgument("*dto.BookReq")).Return(nil, err)
		opts := &server.RouterOpts{
			BookService: mockBService,
		}
		body := testutils.MakeRequestBody(bookPayload)

		req, _ := http.NewRequest("POST", "/books", body)
		_, rec := testutils.ServeReq(opts, req)
		var actualResBody httperror.AppError
		_ = json.Unmarshal(rec.Body.Bytes(), &actualResBody)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, expectedErrBody.Code, actualResBody.Code)
		assert.Equal(t, expectedErrBody.StatusCode, actualResBody.StatusCode)
		assert.NotEmpty(t, actualResBody.Message)
	})
}
