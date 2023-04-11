package services

import (
	"libraryapp/features/users"
	"libraryapp/helper"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	data users.UserData
	vld  *validator.Validate
}

func New(repo users.UserData) users.UserService {
	return &userService{
		data: repo,
		vld:  validator.New(),
	}
}

// Register implements users.UserService
func (us *userService) Register(newUser users.Core) error {
	// Check input validation
	errVld := us.vld.Struct(newUser)
	if errVld != nil {
		return errVld
	}
	// Bcrypt password before insert into database
	passBcrypt, errBcrypt := helper.PassBcrypt(newUser.Password)
	if errBcrypt != nil {
		return errBcrypt
	}
	newUser.Password = passBcrypt
	err := us.data.Register(newUser)
	if err != nil {
		return err
	}
	return nil
}
