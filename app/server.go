package app

import (
	"github.com/gorilla/mux"
	"github.com/mahjadan/login-with-open-telemetry/app/handle"
	"net/http"
	"time"
)

type Server struct {
	http.Server
}

var ServiceName = ""

func New(port string) Server {
	return Server{
		http.Server{
			Addr:         ":" + port,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) SetupRoutes(router *mux.Router, handler handle.Handler) {
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	s.Handler = router
}
