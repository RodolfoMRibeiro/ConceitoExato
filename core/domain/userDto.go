package domain

type UserDto struct {
	Fullname string `json:"fullname"`
	Login    string `json:"login"`
	Email    string `json:"email"`
}
