package router

import (
	_userData "libraryapp/features/users/data"
	_userHandler "libraryapp/features/users/handler"
	_userService "libraryapp/features/users/services"

	_bookData "libraryapp/features/books/data"
	_bookHandler "libraryapp/features/books/handler"
	_bookService "libraryapp/features/books/services"

	_rentData "libraryapp/features/rents/data"
	_rentHandler "libraryapp/features/rents/handler"
	_rentService "libraryapp/features/rents/services"
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
	e.GET("/mybook", bookHdl.MyBook, middlewares.JWTMiddleware())
	e.GET("/booksbyid/:id", bookHdl.GetBookById)
	e.DELETE("/deletebooks/:id", bookHdl.DeleteBook, middlewares.JWTMiddleware())

	rentData := _rentData.New(db)
	rentSrv := _rentService.New(rentData)
	rentHdl := _rentHandler.New(rentSrv)
	e.POST("/rents", rentHdl.Add, middlewares.JWTMiddleware())

}
