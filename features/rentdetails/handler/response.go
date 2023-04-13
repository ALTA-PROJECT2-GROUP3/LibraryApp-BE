package handler

import (
	_book "libraryapp/features/books/handler"
)

type RentdetailResponse struct {
	BookID uint                  `json:"book_id"`
	RentID uint                  `json:"rent_id"`
	Book   _book.AllBookResponse `json:"book"`
}

type ListRentdetailResponse []RentdetailResponse
