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

// HistoryMyBookRented implements rents.RentData
func (rn *rentQuery) HistoryMyBookRented(userID uint) ([]rents.Core, error) {
	tmp := []Rent{}
	tx := rn.db.Preload("RentDetails").Preload("RentDetails.Book").
		Select("rents.*").
		Joins("inner join rentdetails on rents.id = rentdetails.rent_id").
		Joins("inner join books on books.id = rentdetails.book_id").
		Where("books.user_id = ?", userID).Find(&tmp)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat select rent")
		return []rents.Core{}, errors.New("rent not found")
	}
	if tx.Error != nil {
		log.Error("Rent tidak ditemukan")
		return []rents.Core{}, tx.Error
	}

	return ListModelToCore(tmp), nil
}

// History implements rents.RentData
func (rn *rentQuery) HistoryByUserId(userID int) ([]rents.Core, error) {
	tmp := []Rent{}
	tx := rn.db.Preload("RentDetails").Preload("RentDetails.Book").Where("user_id = ?", userID).Find(&tmp)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat select rent")
		return []rents.Core{}, errors.New("rent not found")
	}
	if tx.Error != nil {
		log.Error("Rent tidak ditemukan")
		return []rents.Core{}, tx.Error
	}

	return ListModelToCore(tmp), nil
}

// SelectById implements rents.RentData
func (rn *rentQuery) SelectById(id uint) (rents.Core, error) {
	tmp := Rent{}
	tx := rn.db.Preload("RentDetails").Preload("RentDetails.Book").Where("id = ?", id).First(&tmp)
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
func (rn *rentQuery) Insert(input rents.Core) (uint, error) {
	data := CoreToRent(input)
	tx := rn.db.Create(&data)
	if tx.Error != nil {
		log.Error("Terjadi error saat create")
		return 0, tx.Error
	}
	return data.ID, nil
}
