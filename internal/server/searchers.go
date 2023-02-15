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

// дописать поисковики
// func (s *SearchAuth) Search() []*model.User {
// }

type SearchUsers struct {
	WithUserId int
	WithEmail  string
	WithCityId int
	DB         *sql.DB
}

// поправить rows.Scan (users)
func (s *SearchUsers) Execute() []*model.User {
	users := []*model.User{}
	switch {
	case s.WithUserId != 0:
		rows, err := s.DB.Query("select * from users where user_id=$1;", s.WithUserId)
		if err != nil {
			ErrorLog("Error executing query")
		}
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.Id, &user.Login, &user.Name)
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
			err = rows.Scan(&user.Id, &user.Login, &user.Name)
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
			err = rows.Scan(&user.Id, &user.Login, &user.Name)
			if err != nil {
				ErrorLog("Error scanning rows")
			}
			users = append(users, user)
		}
	}
	return users
}
