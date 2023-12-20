package models

type UserEntity struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phoneNumber"`
}
