package books

import "mime/multipart"

type Core struct {
	Id          uint
	Title       string `validate:"required"`
	Publisher   string `validate:"required"`
	Year        string `validate:"required"`
	Category    string `validate:"required"`
	Description string `validate:"required"`
	Pictures    string
	UserID      uint
	UserName    string
}

type BookService interface {
	Add(newBook Core, file *multipart.FileHeader) error
	GetAll(page int, name string) ([]Core, error)
	Update(userid int, id int, updateBook Core, file *multipart.FileHeader) error
	MyBook(userid int, page int) ([]Core, error)
	GetBookById(id int) (Core, error)
	DeleteBook(userid int, id int) error
}

type BookData interface {
	Insert(input Core) error
	SelectAll(limit, offset int, name string) ([]Core, error)
	Update(userid uint, id uint, input Core) error
	MyBook(userid int, limit, offset int) ([]Core, error)
	GetBookById(id uint) (Core, error)
	DeleteBook(userid int, id int) error
}
