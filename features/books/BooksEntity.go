package books

type Core struct {
	Id          uint
	Title       string `validate:"required"`
	Publisher   string `validate:"required"`
	Year        string `validate:"required"`
	Description string `validate:"required"`
	Pictures    string
	UserID      uint
	UserName    string
}

type BookService interface {
	Add(newBook Core) error
	GetAll(page int, name string) ([]Core, error)
	Update(userid int, id int, updatebook Core)
}

type BookData interface {
	Insert(input Core) error
	SelectAll(limit, offset int, name string) ([]Core, error)
	Update(userid uint, id int, input Core) error
}
