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

// Insert implements rents.RentData
func (rn *rentQuery) Insert(input rents.Core) error {
	data := CoreToRent(input)
	tx := rn.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
