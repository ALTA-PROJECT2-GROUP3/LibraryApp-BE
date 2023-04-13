package handler

import (
	_book "libraryapp/features/books/handler"
	"libraryapp/features/rentdetails/handler"
	"libraryapp/features/rents"
)

type RentResponse struct {
	StartDate  string                       `json:"start_date"`
	EndDate    string                       `json:"end_date"`
	UserID     uint                         `json:"user_id"`
	RentDetail []handler.RentdetailResponse `json:"rent_detail"`
}

type HistoryRentResponse struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ListHistoryResponse []HistoryRentResponse

func CoreToResponse(data rents.Core) RentResponse {
	result := RentResponse{
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
	}

	for _, v := range data.Rentdetail {
		book := _book.AllBookResponse{
			ID:       v.Book.Id,
			Title:    v.Book.Title,
			Pictures: v.Book.Pictures,
			UserName: v.Book.UserName,
		}
		coredetail := handler.RentdetailResponse{
			BookID: v.BookID,
			RentID: v.RentID,
			Book:   book,
		}
		result.RentDetail = append(result.RentDetail, coredetail)

	}

	return result
}

func ListCoreToRoomResp(data []rents.Core) []RentResponse {
	res := []RentResponse{}
	for _, val := range data {
		res = append(res, CoreToResponse(val))
	}
	return res
}
