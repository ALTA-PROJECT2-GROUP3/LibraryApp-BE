package data

import (
	"libraryapp/features/books"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Publisher   string
	Year        string
	Description string
	Pictures    string
	UserID      uint
	UserName    string
}

func CoreToBook(data books.Core) Book {
	return Book{
		Model:       gorm.Model{ID: data.Id},
		Title:       data.Title,
		Publisher:   data.Publisher,
		Year:        data.Year,
		Description: data.Description,
		Pictures:    data.Pictures,
		UserID:      data.UserID,
		UserName:    data.UserName,
	}
}

func BookToCore(data Book) books.Core {
	return books.Core{
		Id:          data.ID,
		Title:       data.Title,
		Publisher:   data.Publisher,
		Year:        data.Year,
		Description: data.Description,
		Pictures:    data.Pictures,
		UserID:      data.UserID,
		UserName:    data.UserName,
	}
}

func ListModelToCore(dataModel []Book) []books.Core {
	var dataCore []books.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, BookToCore(v))
	}
	return dataCore
}
