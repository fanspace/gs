package model

type User struct {
	Id       int64  `json:"id" `
	Account  string `json:"account"`
	Showname string `json:"showname"`
	Email    string `json:"email"`
}
