package handler

type LoginResponse struct {
	Username string `json:"user_name"`
	Name     string `json:"name"`
}

type UserResponse struct {
	// Pictures string `json:"pictures" form:"pictures"`
	Name     string `json:"name" form:"name"`
	Username string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}
