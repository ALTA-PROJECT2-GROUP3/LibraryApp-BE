package services

import (
	"libraryapp/features/books"
	"libraryapp/helper"
	"mime/multipart"

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

// MyBook implements books.BookService
func (srv *bookService) MyBook(userid int, page int) ([]books.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := srv.data.MyBook(userid, limit, offset)
	return data, err
}

// Update implements books.BookService
func (srv *bookService) Update(userid int, id int, updateBook books.Core, fileHeader *multipart.FileHeader) error {
	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/books")
		if err != nil {
			return err
		}
		updateBook.Pictures = uploadURL[0]
	}

	errUpdate := srv.data.Update(uint(userid), uint(id), updateBook)
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
func (srv *bookService) Add(newBook books.Core, fileHeader *multipart.FileHeader) error {
	errValidate := srv.vld.Struct(newBook)
	if errValidate != nil {
		return errValidate
	}

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/books")
		if err != nil {
			return err
		}
		newBook.Pictures = uploadURL[0]
	}

	errInsert := srv.data.Insert(newBook)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (srv *bookService) GetBookById(id int) (books.Core, error) {
	data, err := srv.data.GetBookById(uint(id))
	return data, err
}
func (srv *bookService) DeleteBook(userid int, id int) error {
	errDelete := srv.data.DeleteBook(userid, id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
