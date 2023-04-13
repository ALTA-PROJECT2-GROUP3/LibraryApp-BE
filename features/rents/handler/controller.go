package handler

import (
	"libraryapp/features/rentdetails"
	"libraryapp/features/rents"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	srv       rents.RentService
	srvdetail rentdetails.RentdetailService
}

func New(service rents.RentService, services rentdetails.RentdetailService) *RentHandler {
	return &RentHandler{
		srv:       service,
		srvdetail: services,
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

	rentID, err := rn.srv.Create(newRent)
	if err != nil {
		c.Logger().Error("terjadi kesalahan saat create", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}

	for _, v := range addInput.BookID {
		detail := rentdetails.Core{
			RentID: rentID,
			BookID: uint(v),
		}
		rn.srvdetail.Create(detail)
	}

	return c.JSON(helper.SuccessResponse(http.StatusCreated, "create rent successfully", rentID))
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
	res := CoreToResponse(data)

	return c.JSON(helper.SuccessResponse(http.StatusOK, "detail rent successfully displayed", res))
}

func (rn *RentHandler) History(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	data, err := rn.srv.History(userID)
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	res := ListHistoryResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "success show history", res))
}
