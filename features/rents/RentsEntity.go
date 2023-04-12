package rents

type Core struct {
	Id        uint
	StartDate string `validate:"required"`
	EndDate   string `validate:"required"`
	UserID    uint
	BookID    uint `validate:"required"`
}

type RentService interface {
	Create(newRent Core) error
}

type RentData interface {
	Insert(input Core) error
}
