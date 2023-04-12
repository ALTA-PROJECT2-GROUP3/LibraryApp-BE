package data

import (
	"libraryapp/features/rents"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	StartDate string
	EndDate   string
	UserID    uint
	BookID    uint
}

func CoreToRent(data rents.Core) Rent {
	return Rent{
		Model:     gorm.Model{ID: data.Id},
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
		BookID:    data.BookID,
	}
}

func RentToCore(data Rent) rents.Core {
	return rents.Core{
		Id:        data.ID,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
		BookID:    data.BookID,
	}
}

func ListModelToCore(dataModel []Rent) []rents.Core {
	var dataCore []rents.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, RentToCore(v))
	}
	return dataCore
}
