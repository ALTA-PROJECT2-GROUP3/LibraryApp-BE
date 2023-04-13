package handler

import (
	"libraryapp/features/users"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"

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
		c.Logger().Error("terjadi kesalahan Bind", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	newUser := users.Core{}
	copier.Copy(&newUser, &registerInput)
	err := uh.srv.Register(newUser)
	if err != nil {
		c.Logger().Error("terjadi kesalahan saat register")
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
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}

	file, _ := c.FormFile("pictures")

	updateUser := users.Core{}
	copier.Copy(&updateUser, &updateInput)
	err := uh.srv.Update(userID, updateUser, file)
	if err != nil {
		c.Logger().Error("terjadi kesalahan Saat update", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "UpdateProfile successfully"))

}
func (uh *UserHandler) Profile(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	data, err := uh.srv.Profile(userID)
	if err != nil {
		c.Logger().Error("user tidak ditemukan", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	res := UserResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "profile successfully displayed", res))
}
