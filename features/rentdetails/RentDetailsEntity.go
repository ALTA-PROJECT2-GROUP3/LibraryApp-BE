package rentdetails

import "libraryapp/features/books"

type Core struct {
	Id     uint
	BookID uint `validate:"required"`
	RentID uint `validate:"required"`
	Book   books.Core
}

type RentdetailService interface {
	Create(newRent Core) error
	GetByDetailId(id int) (Core, error)
}

type RentdetailData interface {
	Insert(input Core) error
	SelectByDetailId(id uint) (Core, error)
}
