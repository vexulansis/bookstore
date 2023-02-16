package model

type AuthRespond struct {
	Valid   bool
	Exist   bool
	Message string
}

type SearchRespond struct {
}
type LibraryRespond struct {
	Users []*User
}
