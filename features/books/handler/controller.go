package handler

import (
	"libraryapp/features/books"
	"libraryapp/helper"
	"libraryapp/middlewares"
	"net/http"
	"strconv"

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

func (bk *BookHandler) GetAll(c echo.Context) error {
	var pageNumber int = 1
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam)
		if errConv != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response("Failed, page must number"))
		} else {
			pageNumber = pageConv
		}
	}

	nameParam := c.QueryParam("name")
	data, err := bk.srv.GetAll(pageNumber, nameParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := CoreToGetAllBookResp(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}

func (bk *BookHandler) Update(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	bookID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	// roomID := int(middlewares.ExtractToken(c))
	updateInput := BookRequest{}
	if err := c.Bind(&updateInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}

	updateBook := books.Core{}
	copier.Copy(&updateBook, &updateInput)
	err := bk.srv.Update(userID, bookID, updateBook)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusOK, "update Book successfully"))
}

func (bk *BookHandler) MyBook(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	var pageNumber int = 1
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam)
		if errConv != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response("Failed, page must number"))
		} else {
			pageNumber = pageConv
		}
	}

	// nameParam := c.QueryParam("name")
	data, err := bk.srv.MyBook(userID, pageNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := CoreToGetAllBookResp(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}
func (bk *BookHandler) GetBookById(c echo.Context) error {
	roomID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	data, err := bk.srv.GetBookById(roomID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := MyBookResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "detail book successfully displayed", res))
}
