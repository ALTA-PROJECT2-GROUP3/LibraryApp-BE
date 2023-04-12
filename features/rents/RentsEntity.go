package rents

type Core struct {
	Id        uint
	StartDate string `validate:"required"`
	EndDate   string `validate:"required"`
	UserID    uint
	BookID    string `validate:"required"`
}

type RentService interface {
	Create(newRent Core) (Core, error)
}

type RentData interface {
	Insert(input Core) (Core, error)
}
