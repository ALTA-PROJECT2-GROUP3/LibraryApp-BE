package users

import "mime/multipart"

type Core struct {
	Id       uint
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Address  string `validate:"required"`
	Phone    string `validate:"required"`
	Password string `validate:"required,min=3"`
	Pictures string
}

type UserService interface {
	Register(newUser Core, file *multipart.FileHeader) error
}

type UserData interface {
	Register(newUser Core) error
}
