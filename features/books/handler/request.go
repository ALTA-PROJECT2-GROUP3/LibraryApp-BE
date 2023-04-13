package handler

type BookRequest struct {
	Title       string `json:"title" form:"title"`
	Publisher   string `json:"publisher" form:"publisher"`
	Year        string `json:"year" form:"year"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Pictures    string `json:"pictures" form:"pictures"`
	UserID      uint   `json:"user_id" form:"user_id"`
}
