package handler

import (
	"libraryapp/features/users"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UserService
}

func New(service users.UserService) *UserHandler {
	return &UserHandler{
		srv: service,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	registerInput := RegisterRequest{}
	if err := c.Bind(&registerInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	newUser := users.Core{}
	copier.Copy(&newUser, &registerInput)
	err := uh.srv.Register(newUser)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "register successfully"))
}

func (uh *UserHandler) Login(c echo.Context) error {
	loginInput := LoginRequest{}
	if err := c.Bind(&loginInput); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	token, data, err := uh.srv.Login(loginInput.Username, loginInput.Password)
	if err != nil {
		c.Logger().Error("terjadi kesalahan ")
		return c.JSON(helper.ErrorResponse(err))
	}
	res := LoginResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "login successfully", res, token))
}

func (uh *UserHandler) Update(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	updateInput := UpdateRequest{}
	if err := c.Bind(&updateInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	updateUser := users.Core{}
	copier.Copy(&updateUser, &updateInput)
	err := uh.srv.Update(userID, updateUser)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "UpdateProfile successfully"))

}
func (uh *UserHandler) UserByID(c echo.Context) error {
	userID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	data, err := uh.srv.UserByID(userID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := users.Core{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "profile successfully displayed", res))
}
