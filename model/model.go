package model

type User struct {
	Id       int    `form:"id"`
	Level    string `form:"level"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
