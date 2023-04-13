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

	_rentdetailData "libraryapp/features/rentdetails/data"
	_rentdetailHandler "libraryapp/features/rentdetails/handler"
	_rentdetailService "libraryapp/features/rentdetails/services"
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
	e.GET("/books/:id", bookHdl.GetBookById)
	e.DELETE("/books/:id", bookHdl.DeleteBook, middlewares.JWTMiddleware())

	rentdetailData := _rentdetailData.New(db)
	rentdetailSrv := _rentdetailService.New(rentdetailData)
	rentdetailHdl := _rentdetailHandler.New(rentdetailSrv)
	e.POST("/rentdetails", rentdetailHdl.Add)
	e.GET("/rentdetails/:id", rentdetailHdl.GetById)

	rentData := _rentData.New(db)
	rentSrv := _rentService.New(rentData)
	rentHdl := _rentHandler.New(rentSrv, rentdetailSrv)
	e.POST("/rents", rentHdl.Add, middlewares.JWTMiddleware())
	e.GET("/rents/:id", rentHdl.GetById)
	e.GET("/history", rentHdl.History, middlewares.JWTMiddleware())

}
