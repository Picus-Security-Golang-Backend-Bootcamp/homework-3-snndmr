package book

import (
	"errors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-snndmr/internal/constants"
	"gorm.io/gorm"
	"strings"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Migration To create table.
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Book{})
	if err != nil {
		return
	}
}

// InsertSampleData To insert unique data to database.
func (r *Repository) InsertSampleData(books chan Book) {
	for book := range books {
		r.db.Where(Book{Title: book.Title}).FirstOrCreate(&book)
	}
}

// List To get book list without author data.
func (r *Repository) List() []Book {
	var books []Book
	r.db.Raw("SELECT id, title, stock_id, isbn, page_count, price, stock_count FROM books WHERE is_deleted = ?", false).Find(&books)
	return books
}

// GetById To get book by id.
func (r *Repository) GetById(id int) (error, Book) {
	var book Book
	result := r.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return constants.ErrBookNotFound, Book{}
	}
	return nil, book
}

// Search To search in database with given substr.
func (r *Repository) Search(substr string) []Book {
	lowerSubstr := strings.ToLower(substr)
	var books []Book
	r.db.Where("is_deleted = ? AND (title LIKE ? OR isbn LIKE ? OR author LIKE ?)", false, "%"+lowerSubstr+"%", "%"+lowerSubstr+"%", "%"+lowerSubstr+"%").Find(&books)
	return books
}

// Update To update book after delete and buy operations.
func (r *Repository) Update(book Book) error {
	result := r.db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
