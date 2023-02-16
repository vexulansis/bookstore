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
		authRespond := s.SignUp(user)
		err = tmpl.Execute(w, authRespond)
		if err != nil {
			ErrorLog("Error executing form.html template")
		}
	}
}
