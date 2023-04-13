package rentdetails

type Core struct {
	Id     uint
	BookID uint `validate:"required"`
	RentID uint `validate:"required"`
}

type RentdetailService interface {
	Create(newRent Core) error
	GetByDetailId(id int) (Core, error)
}

type RentdetailData interface {
	Insert(input Core) error
	SelectByDetailId(id uint) (Core, error)
}
