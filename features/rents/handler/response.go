package handler

type RentResponse struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	UserID    string `json:"user_id" form:"user_id"`
}

type HistoryRentResponse struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

type ListHistoryResponse []HistoryRentResponse
