package book

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-snndmr/internal/constants"
)

type Deletable interface {
	Delete() error
}

type Book struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"type:varchar(100)"`
	StockId    int
	ISBN       string `gorm:"type:varchar(20)"`
	PageCount  int
	Price      float32
	StockCount int
	Author     string `gorm:"type:varchar(100)"`
	IsDeleted  bool
}

// Delete To delete book.
func (book *Book) Delete() error {
	if book.IsDeleted {
		return constants.ErrBookAlreadyDeleted
	}

	book.IsDeleted = true
	return nil
}

// DecreaseAmount To buy book.
func (book *Book) DecreaseAmount(amount int) error {
	if amount < 0 {
		return constants.ErrNegativeAmount
	}

	if amount > book.StockCount {
		return constants.ErrBookOutOfStock
	}

	book.StockCount -= amount
	return nil
}

func (book *Book) ToString() string {
	return fmt.Sprintf("ID: %3d | Title: %-25s | Author: %-25s | ISBN: %-10s | StockCode: %5d | Price: ₺%.2f | Stock Count: %3d",
		book.ID,
		book.Title,
		book.Author,
		book.ISBN,
		book.StockId,
		book.Price,
		book.StockCount)
}

func (book *Book) ToStringWithoutAuthor() string {
	return fmt.Sprintf("ID: %3d | Title: %-25s | ISBN: %-10s | StockCode: %5d | Price: ₺%.2f | Stock Count: %3d",
		book.ID,
		book.Title,
		book.ISBN,
		book.StockId,
		book.Price,
		book.StockCount)
}
