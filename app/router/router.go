package router

import (
	_userData "libraryapp/features/users/data"
	_userHandler "libraryapp/features/users/handler"
	_userService "libraryapp/features/users/services"

	_bookData "libraryapp/features/books/data"
	_bookHandler "libraryapp/features/books/handler"
	_bookService "libraryapp/features/books/services"
	"libraryapp/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userSrv := _userService.New(userData)
	userHdl := _userHandler.New(userSrv)
	e.POST("/register", userHdl.Register)
	e.POST("/login", userHdl.Login)
	e.PUT("/users", userHdl.Update, middlewares.JWTMiddleware())
	e.GET("/users", userHdl.Profile, middlewares.JWTMiddleware())

	bookData := _bookData.New(db)
	bookSrv := _bookService.New(bookData)
	bookHdl := _bookHandler.New(bookSrv)
	e.POST("/books", bookHdl.Add, middlewares.JWTMiddleware())
	e.GET("/books", bookHdl.GetAll)
	e.PUT("/books/:id", bookHdl.Update, middlewares.JWTMiddleware())
	e.GET("/booksbyid/:id", bookHdl.GetBookById)
}
