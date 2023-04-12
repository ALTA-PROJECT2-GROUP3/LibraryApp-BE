package handler

type RentResponse struct {
	ID        uint   `json:"id"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	UserID    uint   `json:"user_id" form:"user_id"`
	BookID    uint   `json:"book_id" form:"book_id"`
}
