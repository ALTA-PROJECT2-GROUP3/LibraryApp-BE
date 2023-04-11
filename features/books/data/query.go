package data

import (
	"libraryapp/features/books"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) books.BookData {
	return &bookQuery{
		db: db,
	}
}

// SelectAll implements books.BookData
func (bk *bookQuery) SelectAll(limit int, offset int, name string) ([]books.Core, error) {
	nameSearch := "%" + name + "%"
	var roomsModel []Book
	tx := bk.db.Limit(limit).Offset(offset).Where("rooms.name LIKE ?", nameSearch).Select("rooms.id, rooms.name, rooms.price, rooms.pictures, rooms.availability, rooms.description, rooms.location, users.name AS user_name, AVG(feedbacks.rating) AS rating_room").Joins("JOIN users ON rooms.user_id = users.id").Joins("LEFT JOIN feedbacks ON rooms.id = feedbacks.room_id").Group("rooms.id").Find(&roomsModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	roomsCoreAll := ListModelToCore(roomsModel)
	return roomsCoreAll, nil
}

// Insert implements books.BookData
func (bk *bookQuery) Insert(input books.Core) error {
	data := CoreToBook(input)
	tx := bk.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
