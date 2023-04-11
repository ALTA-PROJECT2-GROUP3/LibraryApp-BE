package router

import (
	_userData "libraryapp/features/users/data"
	_userHandler "libraryapp/features/users/handler"
	_userService "libraryapp/features/users/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userSrv := _userService.New(userData)
	userHdl := _userHandler.New(userSrv)
	e.POST("/register", userHdl.Register)
	e.POST("/login", userHdl.Login)
}
