package handler

import (
	"git.garena.com/sea-labs-id/trainers/library-api/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostBorrowingRecord(c *gin.Context) {
	payload, _ := c.Get("payload")
	borrowingReq, _ := payload.(*dto.BorrowingReq)
	books, err := h.borrowingService.BorrowBook(borrowingReq)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) PatchBorrowingRecord(c *gin.Context) {
	recordID, _ := strconv.Atoi(c.Param("id"))
	books, err := h.borrowingService.ReturnBook(recordID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, books)
}
