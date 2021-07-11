package model

type UserDetail struct {
	UserId      string
	UserName    string
	Password    string
	Gender      string
	Email       string
	Phone       string
	Authorities []string
}
