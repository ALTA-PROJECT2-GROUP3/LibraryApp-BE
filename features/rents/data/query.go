package data

import (
	"errors"
	"libraryapp/features/rents"

	"github.com/labstack/gommon/log"
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

// SelectById implements rents.RentData
func (rn *rentQuery) SelectById(id uint) (rents.Core, error) {
	tmp := Rent{}
	tx := rn.db.Where("id = ?", id).First(&tmp)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat select rent")
		return rents.Core{}, errors.New("rent not found")
	}
	if tx.Error != nil {
		log.Error("Rent tidak ditemukan")
		return rents.Core{}, tx.Error
	}

	return RentToCore(tmp), nil
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
