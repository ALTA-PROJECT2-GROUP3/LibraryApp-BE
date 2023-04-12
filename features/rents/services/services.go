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

// History implements rents.RentService
func (srv *rentService) History(userID int) ([]rents.Core, error) {
	tmp, err := srv.data.History(userID)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// GetById implements rents.RentService
func (srv *rentService) GetById(id int) (rents.Core, error) {
	data, err := srv.data.SelectById(uint(id))
	return data, err
}

// Create implements rents.RentService
func (srv *rentService) Create(newRent rents.Core) error {
	errValidate := srv.vld.Struct(newRent)
	if errValidate != nil {
		return errValidate
	}

	errInsert := srv.data.Insert(newRent)
	if errInsert != nil {
		return errInsert
	}
	return nil
}
