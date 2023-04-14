package services

import (
	"libraryapp/features/rentdetails"

	"github.com/go-playground/validator/v10"
)

type rentdetailsService struct {
	data rentdetails.RentdetailData
	vld  *validator.Validate
}

func New(repo rentdetails.RentdetailData) rentdetails.RentdetailService {
	return &rentdetailsService{
		data: repo,
		vld:  validator.New(),
	}
}

// Create implements rentdetails.RentdetailService
func (srv *rentdetailsService) Create(newRent rentdetails.Core) error {
	errValidate := srv.vld.StructExcept(newRent, "Book")
	if errValidate != nil {
		return errValidate
	}

	errInsert := srv.data.Insert(newRent)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// GetByDetailId implements rentdetails.RentdetailService
func (srv *rentdetailsService) GetByDetailId(id int) (rentdetails.Core, error) {
	data, err := srv.data.SelectByDetailId(uint(id))
	return data, err
}
