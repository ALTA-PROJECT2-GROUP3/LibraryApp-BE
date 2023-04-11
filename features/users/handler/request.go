package handler

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}
type LoginRequest struct {
	Username string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	// Pictures string `json:"pictures" form:"pictures"`
	Name     string `json:"name" form:"name"`
	Username string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}
