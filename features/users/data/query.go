package data

import (
	"errors"
	"libraryapp/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

// Login implements users.UserData
func (uq *userQuery) Login(username string) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("username = ?", username).First(&tmp)
	if tx.RowsAffected < 1 {
		return users.Core{}, errors.New("username not found")
	}
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}

// Register implements users.UserData
func (uq *userQuery) Register(newUser users.Core) error {
	data := CoreToUser(newUser)
	tx := uq.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (uq *userQuery) Update(userID int, updateUser users.Core) error {
	update := CoreToUser(updateUser)
	tx := uq.db.Model(&User{}).Where("id = ?", userID).Updates(&update)
	if tx.RowsAffected < 1 {
		return errors.New("profile no updated")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (uq *userQuery) UserByID(userID int) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("id = ?", userID).First(&tmp)
	if tx.RowsAffected < 1 {
		return users.Core{}, errors.New("user not found")
	}
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}
