package handler

import (
	"libraryapp/features/books"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	srv books.BookService
}

func New(service books.BookService) *BookHandler {
	return &BookHandler{
		srv: service,
	}
}

func (bk *BookHandler) Add(c echo.Context) error {
	addInput := BookRequest{}
	addInput.UserID = uint(middlewares.ExtractToken(c))
	if err := c.Bind(&addInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}

	newBook := books.Core{}
	copier.Copy(&newBook, &addInput)

	err := bk.srv.Add(newBook)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "add book successfully"))
}
