package data

import (
	"libraryapp/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Phone    string
	Pictures string
}

func CoreToUser(data users.Core) User {
	return User{
		Model:    gorm.Model{ID: data.Id},
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
		Phone:    data.Phone,
		Pictures: data.Pictures,
	}
}

func UserToCore(data User) users.Core {
	return users.Core{
		Id:       data.ID,
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
		Phone:    data.Phone,
		Pictures: data.Pictures,
	}
}
