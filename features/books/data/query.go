package data

import (
	"libraryapp/features/books"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) books.BookData {
	return &bookQuery{
		db: db,
	}
}

// Insert implements books.BookData
func (bk *bookQuery) Insert(input books.Core) error {
	data := CoreToBook(input)
	tx := bk.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
