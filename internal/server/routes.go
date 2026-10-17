package server

import (
	"myclockspace.kevinmartiniano.com/internal/app"
	"myclockspace.kevinmartiniano.com/internal/routes"
)

type Routes struct {
	register map[string]func(r *app.Router)
	router   *app.Router
}

func NewRoutes(r *app.Router) *Routes {
	return &Routes{
		register: register(r),
		router:   r,
	}
}

func (r *Routes) RegisterRoutes() {
	for _, route := range r.register {
		route(r.router)
	}
}

func register(r *app.Router) map[string]func(r *app.Router) {
	return map[string]func(r *app.Router){
		"healthcheck": func(r *app.Router) { routes.NewHealthcheck().Register(r) },
	}
}
