package dto

type BookReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Cover       string `json:"cover"`
	AuthorID    uint   `json:"authorID" binding:"required"`
}
