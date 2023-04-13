package data

import (
	"errors"
	"libraryapp/features/rentdetails"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type rentdetailQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) rentdetails.RentdetailData {
	return &rentdetailQuery{
		db: db,
	}
}

// Insert implements rentdetails.RentdetailData
func (rn *rentdetailQuery) Insert(input rentdetails.Core) error {
	data := CoreToRentdetail(input)
	tx := rn.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectByDetailId implements rentdetails.RentdetailData
func (rn *rentdetailQuery) SelectByDetailId(id uint) (rentdetails.Core, error) {
	tmp := Rentdetail{}
	tx := rn.db.Where("id = ?", id).First(&tmp)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat select rent")
		return rentdetails.Core{}, errors.New("rent not found")
	}
	if tx.Error != nil {
		log.Error("Rent tidak ditemukan")
		return rentdetails.Core{}, tx.Error
	}

	return RentdetailToCore(tmp), nil
}
