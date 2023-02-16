package server

import (
	"database/sql"
	. "project/internal"
	"project/internal/model"
)

type Searcher interface {
	Search()
}
type SearchAuth struct {
	WithUserId   int
	WithLogin    string
	WithPassword string
	DB           *sql.DB
}

func (s *SearchAuth) Search() []*model.User {
	users := []*model.User{}
	switch {
	case s.WithUserId != 0:
		rows, err := s.DB.Query("select * from auth where user_id=$1;", s.WithUserId)
		if err != nil {
			ErrorLog("Error executing query")
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.User_id, &user.Login, &user.Password)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	case s.WithLogin != "":
		rows, err := s.DB.Query("select * from auth where login=$1;", s.WithLogin)
		if err != nil {
			ErrorLog("Error executing query")
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.User_id, &user.Login, &user.Password)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	case s.WithPassword != "":
		rows, err := s.DB.Query("select * from auth where password=$1;", s.WithPassword)
		if err != nil {
			ErrorLog("Error executing query")
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.User_id, &user.Login, &user.Password)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	}
	return users
}

type SearchUsers struct {
	WithUserId int
	WithEmail  string
	WithCityId int
	DB         *sql.DB
}

func (s *SearchUsers) Search() []*model.User {
	users := []*model.User{}
	switch {
	case s.WithUserId != 0:
		rows, err := s.DB.Query("select * from users where user_id=$1;", s.WithUserId)
		if err != nil {
			ErrorLog("Error executing query")
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.User_id, &user.Email, &user.City_id)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	case s.WithEmail != "":
		rows, err := s.DB.Query("select * from users where email=$1;", s.WithEmail)
		if err != nil {
			ErrorLog(err.Error())
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.User_id, &user.Email, &user.City_id)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	case s.WithCityId != 0:
		rows, err := s.DB.Query("select * from users where city_id=$1", s.WithCityId)
		if err != nil {
			ErrorLog("Error executing query")
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.User_id, &user.Email, &user.City_id)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	}
	return users
}
