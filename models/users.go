package models

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserInfo struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}