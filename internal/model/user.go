package model

type User struct {
	User_id  int    `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	City_id  int    `json:"city_id"`
}
