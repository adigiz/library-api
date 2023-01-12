package handler

import (
	"net/http"

	"library-api/dto"
	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get list of books available in library
// @Description Get list of books available in library
// @Tags books
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /books [get]
func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetBooks()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) PostBook(c *gin.Context) {
	payload, _ := c.Get("payload")
	bookReq, _ := payload.(*dto.BookReq)
	insertedBook, err := h.bookService.AddBook(bookReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, *insertedBook)
}
