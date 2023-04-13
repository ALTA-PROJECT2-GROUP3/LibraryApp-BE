package handler

type RentdetailRequest struct {
	BookID uint `json:"book_id" form:"book_id"`
	RentID uint `json:"rent_id" form:"rent_id"`
}
