package handler

type LoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type ProfileResponse struct {
	ID       uint   `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}
