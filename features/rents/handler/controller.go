package handler

import (
	"libraryapp/features/rents"
)

type RentHandler struct {
	srv rents.RentService
}

func New(service rents.RentService) *RentHandler {
	return &RentHandler{
		srv: service,
	}
}
