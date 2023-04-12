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
// Update implements books.BookService
func (srv *bookService) Update(userid int, id int, updatebook books.Core) {
	errUpdate := srv.data.Update(uint(userid), id, updatebook)
	if errUpdate != nil {
		return errUpdate
	}
	
	errUpdate := srv.data.Update(updatebook)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}



// GetAll implements books.BookService
func (srv *bookService) GetAll(page int, name string) ([]books.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := srv.data.SelectAll(limit, offset, name)
	return data, err
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
