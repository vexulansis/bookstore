package server

import (
	. "project/internal"
	"project/internal/model"
)

func (s *Server) NewUser(u *model.User) {
	tx, err := s.Store.DB.Begin()
	if err != nil {
		ErrorLog("Error starting transaction")
	}
	defer tx.Rollback()
	tx.QueryRow("insert into users default values returning user_id").Scan(&u.User_id)
	_, err = tx.Exec("insert into auth (user_id, login, password) values ($1, $2, $3)", u.User_id, u.Login, u.Password)
	if err != nil {
		ErrorLog(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		ErrorLog(err.Error())
	}
}
func (s *Server) UserAccess(u *model.User) bool {
	su := SearchAuth{
		WithLogin: u.Login,
		DB:        s.Store.DB,
	}
	ref := su.Search()[0]
	if ref.Password == u.Password {
		return true
	}
	return false
}
