package dto

type BorrowingReq struct {
	UserID int `json:"userId" binding:"required"`
	BookID int `json:"bookID" binding:"required"`
}
