package rents

import "libraryapp/features/rentdetails"

type Core struct {
	Id         uint
	StartDate  string `validate:"required"`
	EndDate    string `validate:"required"`
	UserID     uint
	Rentdetail []rentdetails.Core
}

type RentService interface {
	Create(newRent Core) (uint, error)
	GetById(id int) (Core, error)
	HistoryByUserId(userID int) ([]Core, error)
	HistoryMyBookRented(userID uint) ([]Core, error)
}

type RentData interface {
	Insert(input Core) (uint, error)
	SelectById(id uint) (Core, error)
	HistoryByUserId(userID int) ([]Core, error)
	HistoryMyBookRented(userID uint) ([]Core, error)
}
