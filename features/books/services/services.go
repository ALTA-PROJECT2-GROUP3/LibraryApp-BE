package services

import (
	"libraryapp/features/books"

	"github.com/go-playground/validator/v10"
)

type bookService struct {
	data books.BookData
	vld  *validator.Validate
}

func New(repo books.BookData) books.BookService {
	return &bookService{
		data: repo,
		vld:  validator.New(),
	}
}

// Add implements books.BookService
func (srv *bookService) Add(newBook books.Core) error {
	errValidate := srv.vld.Struct(newBook)
	if errValidate != nil {
		return errValidate
	}

	errInsert := srv.data.Insert(newBook)
	if errInsert != nil {
		return errInsert
	}
	return nil
}
