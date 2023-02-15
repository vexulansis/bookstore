package server

import (
	"project/internal/model"
)

func (s *Server) UserExist(u *model.User) bool {
	// su := SearchAuth{
	// 	WithLogin: u.Login,
	// 	DB:        s.Store.DB,
	// }
	// // if len(su.Search()) == 0 {
	// // 	return false
	// // }
	return true
}

func (s *Server) ValidateUser(u *model.User) bool {
	return true
}
