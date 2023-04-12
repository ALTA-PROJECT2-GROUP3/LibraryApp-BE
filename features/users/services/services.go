package services

import (
	"libraryapp/features/users"
	"libraryapp/helper"
	"libraryapp/middlewares"

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

// Login implements users.UserService
func (us *userService) Login(username string, password string) (string, users.Core, error) {
	tmp, err := us.data.Login(username)
	if err != nil {
		return "", users.Core{}, err
	}
	// Compare password
	if err := helper.PassCompare(tmp.Password, password); err != nil {
		return "", users.Core{}, err
	}
	// Generate token
	token, err := middlewares.CreateToken(tmp.Id)
	if err != nil {
		return "", users.Core{}, err
	}
	return token, tmp, nil
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

func (us *userService) Update(userID int, updatedUser users.Core) error {
	// Check input validation
	errVld := us.vld.Struct(updatedUser)
	if errVld != nil {
		return errVld
	}
	// Bcrypt password before updating into database
	if updatedUser.Password != "" {
		passBcrypt, errBcrypt := helper.PassBcrypt(updatedUser.Password)
		if errBcrypt != nil {
			return errBcrypt
		}
		updatedUser.Password = passBcrypt
	}
	err := us.data.Update(userID, updatedUser)
	if err != nil {
		return err
	}
	return nil
}
func (us *userService) UserByID(userID int) (users.Core, error) {
	tmp, err := us.data.UserByID(userID)
	if err != nil {
		return users.Core{}, err
	}
	return tmp, nil
}
