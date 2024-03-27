package handler

type LoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type ProfileResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}
