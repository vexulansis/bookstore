package server

import (
	"html/template"
	"net/http"
	. "project/internal"
	"project/internal/model"

	_ "github.com/lib/pq"
)

func (s *Server) AuthGet(w http.ResponseWriter, r *http.Request) {
	// HTML execution
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		ErrorLog("Error parsing form.html file")
	}
	res := model.NewRespond()
	err = tmpl.Execute(w, res)
	if err != nil {
		ErrorLog("Error executing form.html template")
	}

}
func (s *Server) AuthPost(w http.ResponseWriter, r *http.Request) {
	// User validation
	user := &model.User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	res := s.SignUp(user)
	// HTML execution
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		ErrorLog("Error parsing form.html file")
	}
	err = tmpl.Execute(w, res)
	if err != nil {
		ErrorLog("Error executing form.html template")
	}
}
func (s *Server) SearchGet(w http.ResponseWriter, r *http.Request) {
	// HTML execution
	tmpl, err := template.ParseFiles("interface.html")
	if err != nil {
		ErrorLog("Error parsing interface.html file")
	}
	err = tmpl.Execute(w, 0)
	if err != nil {
		ErrorLog("Error executing interface.html template")
	}
	//
}
func (s *Server) SearchPost(w http.ResponseWriter, r *http.Request) {
	// HTML execution
	tmpl, err := template.ParseFiles("interface.html")
	if err != nil {
		ErrorLog("Error parsing interface.html file")
	}
	err = tmpl.Execute(w, 0)
	if err != nil {
		ErrorLog("Error executing interface.html template")
	}
	// Form parsing
	request := model.RequestUser{
		By:   r.FormValue("search_by"),
		Data: r.FormValue("data"),
	}
	s.SearchUser(&request)
}
func (s *Server) LibraryGet(w http.ResponseWriter, r *http.Request) {
	table := s.GetAllUsers()
	tmpl, err := template.ParseFiles("library.html")
	if err != nil {
		ErrorLog("Error parsing library.html file")
	}
	err = tmpl.Execute(w, table)
}
