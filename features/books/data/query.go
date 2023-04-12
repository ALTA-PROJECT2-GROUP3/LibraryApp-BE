package data

import (
	"libraryapp/features/books"
	"errors"
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

// SelectAll implements books.BookData
func (bk *bookQuery) SelectAll(limit int, offset int, name string) ([]books.Core, error) {
	nameSearch := "%" + name + "%"
	var booksModel []Book
	tx := bk.db.Limit(limit).Offset(offset).Where("books.title LIKE ?", nameSearch).Select("books.id, books.title, books.pictures, users.name AS user_name").Joins("JOIN users ON books.user_id = users.id").Group("books.id").Find(&booksModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	booksCoreAll := ListModelToCore(booksModel)
	return booksCoreAll, nil
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

func (bk *bookQuery) Update(userId uint, input books.Core, id uint) error {
	data := CoreToBook(input)
	tx := bk.db.Model(&books.Core{}).Where("id = ? AND user_id = ?", id, userId).Updates(&data)
	if tx.RowsAffected < 1 {
		return errors.New("book no updated")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
