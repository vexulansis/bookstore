package server

import (
	"net/http"
	. "project/internal"
	"project/internal/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var corsObj = handlers.AllowedOrigins([]string{"*"})

type Server struct {
	Router *mux.Router
	Logger *logrus.Logger
	Store  *store.Store
}

func NewServer(store *store.Store) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		Store:  store,
	}
	s.configureRouter()
	return s
}
func (s *Server) configureRouter() {
	s.Router.Use(handlers.CORS(corsObj))
	s.Router.HandleFunc("/", s.Auth).Methods("GET")
	s.Router.HandleFunc("/", s.Auth).Methods("POST")
	s.Router.HandleFunc("/library", s.Library).Methods("GET")
	s.Router.HandleFunc("/library", s.Library).Methods("POST")
}

func (s *Server) Start() {
	err := http.ListenAndServe(":8080", handlers.CORS(corsObj)(s.Router))
	if err != nil {
		FatalLog("Error listening and serving")
	}
}
