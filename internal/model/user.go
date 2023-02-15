package model

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) Valid() (bool, string) {
	switch {
	case len(u.Login) < 4:
		return false, "Login must be 4+ letters"
	case len(u.Password) < 8:
		return false, "Password must be 8+ letters"
	}
	return true, ""
}

func (u *User) Sanitize() {
	u.Password = ""
}
