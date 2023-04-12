package data

import (
	"libraryapp/features/rents"

	"gorm.io/gorm"
)

type rentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) rents.RentData {
	return &rentQuery{
		db: db,
	}
}
