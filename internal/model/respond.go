package model

type Respond struct {
	UserExist bool
	UserValid bool
	Message   string
}

func NewRespond() *Respond {
	return &Respond{
		UserValid: false,
		Message:   "",
	}
}
