package dto

import "library-api/models"

type BookRes struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	AuthorID    uint   `json:"authorID"`
}

func (_ *BookRes) FromBook(b *models.Book) *BookRes {
	return &BookRes{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Cover:       b.Cover,
		AuthorID:    b.AuthorID,
	}
}
