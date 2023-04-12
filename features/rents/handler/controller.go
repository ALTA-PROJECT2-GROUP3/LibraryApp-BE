package handler

import (
	"libraryapp/features/rents"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"

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
