package services

import (
	"libraryapp/features/users"
	"mime/multipart"

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
func (*userService) Register(newUser users.Core, file *multipart.FileHeader) error {
	panic("unimplemented")
}
