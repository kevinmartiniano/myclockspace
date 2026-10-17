package server

import (
	"net/http"

	"myclockspace.kevinmartiniano.com/internal/app"
)

type Server struct {
	router *app.Router
}

func NewServer(r *app.Router) *Server {
	return &Server{
		router: r,
	}
}

func (s *Server) Up() {
	RegisterRoutes(s.router)

	if err := http.ListenAndServe(":8080", s.router); err != nil {
		panic(err)
	}
}
