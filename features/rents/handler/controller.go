package handler

import (
	"libraryapp/features/rents"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	srv rents.RentService
}

func New(service rents.RentService) *RentHandler {
	return &RentHandler{
		srv: service,
	}
}

func (rn *RentHandler) Add(c echo.Context) error {
	addInput := RentRequest{}
	addInput.UserID = uint(middlewares.ExtractToken(c))
	if err := c.Bind(&addInput); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}

	newRent := rents.Core{}
	copier.Copy(&newRent, &addInput)

	err := rn.srv.Create(newRent)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "create rent successfully"))
}

func (rn *RentHandler) GetById(c echo.Context) error {
	rentID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		c.Logger().Error("terjadi kesalahan")
		return errCnv
	}
	data, err := rn.srv.GetById(rentID)
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	res := RentResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "detail rent successfully displayed", res))
}
