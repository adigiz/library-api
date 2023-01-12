package repository

import (
	"library-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository interface {
	FindBooks() ([]*models.Book, error)
	Save(book *models.Book) (*models.Book, int, error)
	FindBook(id int) (*models.Book, error)
	DecreaseStock(id uint) error
		IncreaseStock(id int) error
}

type bookRepository struct {
	db *gorm.DB
}

type BROpts struct {
	DB *gorm.DB
}

func NewBookRepository(o *BROpts) BookRepository {
	return &bookRepository{
		db: o.DB,
	}
}

func (repo *bookRepository) FindBooks() ([]*models.Book, error) {
	var books []*models.Book
	result := repo.db.Joins("Author").Find(&books)
	return books, result.Error
}

func (repo *bookRepository) Save(book *models.Book) (b *models.Book, rowsEffected int, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(book)
	return book, int(result.RowsAffected), result.Error
}

func (repo *bookRepository) FindBook(id int) (*models.Book, error) {
	var book *models.Book
	result := repo.db.Find(&book, id)
	return book, result.Error
}

func (repo *bookRepository) DecreaseStock(id uint) error {
	result := repo.db.Raw("UPDATE books SET stock = stock - 1 WHERE id = ?", id).Scan(nil)
	return result.Error
}

func (repo *bookRepository) IncreaseStock(id int) error {
	result := repo.db.Raw("UPDATE books SET stock = stock + 1 WHERE id = ?", id).Scan(nil)
	return result.Error
}
