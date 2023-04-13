package handler

import (
	"libraryapp/features/rentdetails"
	"libraryapp/helper"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type RentdetailHandler struct {
	srv rentdetails.RentdetailService
}

func New(service rentdetails.RentdetailService) *RentdetailHandler {
	return &RentdetailHandler{
		srv: service,
	}
}

func (rn *RentdetailHandler) Add(c echo.Context) error {
	addInput := RentdetailRequest{}
	if err := c.Bind(&addInput); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}

	newRent := rentdetails.Core{}
	copier.Copy(&newRent, &addInput)

	err := rn.srv.Create(newRent)
	if err != nil {
		c.Logger().Error("terjadi kesalahan saat create", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "create rent successfully"))
}

func (rn *RentdetailHandler) GetById(c echo.Context) error {
	rentdetailID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		c.Logger().Error("terjadi kesalahan")
		return errCnv
	}
	data, err := rn.srv.GetByDetailId(rentdetailID)
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	res := RentdetailResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "detail rent successfully displayed", res))
}
