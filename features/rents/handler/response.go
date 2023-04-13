package handler

import (
	"libraryapp/features/rentdetails/handler"
	"libraryapp/features/rents"
)

type RentResponse struct {
	StartDate  string                       `json:"start_date" form:"start_date"`
	EndDate    string                       `json:"end_date" form:"end_date"`
	UserID     uint                         `json:"user_id" form:"user_id"`
	RentDetail []handler.RentdetailResponse `json:"rent_detail"`
}

type HistoryRentResponse struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

type ListHistoryResponse []HistoryRentResponse

func CoreToResponse(data rents.Core) RentResponse {
	result := RentResponse{
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
	}

	for _, v := range data.Rentdetail {
		coredetail := handler.RentdetailResponse{
			BookID: v.BookID,
			RentID: v.RentID,
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
