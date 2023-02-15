package server

import (
	"database/sql"
	"fmt"
	. "project/internal"
	"project/internal/model"
)

func (s *Server) SearchUser(r *model.RequestUser) *model.User {
	var rows *sql.Rows
	var err error
	var query string
	switch r.By {
	case "id":
		query = fmt.Sprintf("select * from users where %s=%s", r.By, r.Data)
		rows, err = s.Store.DB.Query(query)
	case "login":
		query = fmt.Sprintf("select * from users where %s='%s'", r.By, r.Data)
		rows, err = s.Store.DB.Query(query)
	}
	if err != nil {
		ErrorLog("Error executing query")
	}
	for rows.Next() {
		user := new(model.User)
		err = rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			ErrorLog("Error scanning rows")
		}
		return user
	}
	return nil
}
func (s *Server) GetAllUsers() *[]model.User {
	users := []model.User{}
	query := "select * from users"
	rows, err := s.Store.DB.Query(query)
	for rows.Next() {
		user := new(model.User)
		err = rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			ErrorLog("Error scanning rows")
		}
		users = append(users, *user)
	}
	return &users
}
