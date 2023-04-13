package handler

import (
	"libraryapp/features/books"
)

type AllBookResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Pictures string `json:"pictures"`
	UserName string `json:"user_name"`
}

func CoreToGetAllBookRespB(data books.Core) AllBookResponse {
	return AllBookResponse{
		ID:       data.Id,
		Title:    data.Title,
		Pictures: data.Pictures,
		UserName: data.UserName,
	}
}

func CoreToGetAllBookResp(data []books.Core) []AllBookResponse {
	res := []AllBookResponse{}
	for _, val := range data {
		res = append(res, CoreToGetAllBookRespB(val))
	}
	return res
}

type MyBookResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Pictures string `json:"pictures"`
}

func CoreToGetMyBook(data books.Core) MyBookResponse {
	return MyBookResponse{
		ID:       data.Id,
		Title:    data.Title,
		Pictures: data.Pictures,
	}
}

func CoreToGetMyBookResp(data []books.Core) []MyBookResponse {
	res := []MyBookResponse{}
	for _, val := range data {
		res = append(res, CoreToGetMyBook(val))
	}
	return res
}
