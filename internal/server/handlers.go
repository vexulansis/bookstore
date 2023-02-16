package server

import (
	"net/http"
	. "project/internal"
	"project/internal/model"
	"text/template"

	_ "github.com/lib/pq"
)

func (s *Server) Auth(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		ErrorLog("Error parsing form.html file")
	}
	switch r.Method {
	case "GET":
		authRespond := model.AuthRespond{
			Valid:   true,
			Exist:   false,
			Message: "",
		}
		err = tmpl.Execute(w, authRespond)
		if err != nil {
			ErrorLog("Error executing form.html template")
		}
	case "POST":
		user := &model.User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		switch r.FormValue("button") {
		case "signup":
			authRespond := s.SignUp(user)
			err = tmpl.Execute(w, authRespond)
			if err != nil {
				ErrorLog("Error executing form.html template")
			}
		case "signin":
			authRespond := s.SignIn(user)
			err = tmpl.Execute(w, authRespond)
			if err != nil {
				ErrorLog("Error executing form.html template")
			}
		}
	}
}
func (s *Server) Library(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("library.html")
	if err != nil {
		ErrorLog("Error parsing library.html file")
	}
	su := SearchAuth{
		DB: s.Store.DB,
	}
	libraryRespond := model.LibraryRespond{
		Users: su.Search(),
	}
	err = tmpl.Execute(w, libraryRespond.Users)
}
