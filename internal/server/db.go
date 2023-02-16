package server

import (
	. "project/internal"
	"project/internal/model"
)

// разобраться с транзакциями

func (s *Server) NewUser(u *model.User) {
	InfoLog(u.Password)
	tx, err := s.Store.DB.Begin()
	if err != nil {
		ErrorLog("Error starting transaction")
	}
	defer tx.Rollback()
	tx.QueryRow("insert into users default values returning user_id").Scan(&u.User_id)
	_, err = tx.Exec("insert into auth (user_id, login, password) values ($1, $2, $3) ", u.User_id, u.Login, u.Password)
	if err != nil {
		ErrorLog(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		ErrorLog(err.Error())
	}
}
