package services

import (
	"libraryapp/features/rents"

	"github.com/go-playground/validator/v10"
)

type rentService struct {
	data rents.RentData
	vld  *validator.Validate
}

func New(repo rents.RentData) rents.RentService {
	return &rentService{
		data: repo,
		vld:  validator.New(),
	}
}
