package data

import (
	_book "libraryapp/features/books/data"
	"libraryapp/features/rentdetails"

	"gorm.io/gorm"
)

type Rentdetail struct {
	gorm.Model
	BookID uint
	RentID uint
	Book   _book.Book `gorm:"foreignKey:BookID"`
}

func CoreToRentdetail(data rentdetails.Core) Rentdetail {
	return Rentdetail{
		Model:  gorm.Model{ID: data.Id},
		BookID: data.BookID,
		RentID: data.RentID,
	}
}

func RentdetailToCore(data Rentdetail) rentdetails.Core {
	return rentdetails.Core{
		Id:     data.ID,
		BookID: data.BookID,
		RentID: data.RentID,
	}
}

func ListModelToCore(dataModel []Rentdetail) []rentdetails.Core {
	var dataCore []rentdetails.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, RentdetailToCore(v))
	}
	return dataCore
}
