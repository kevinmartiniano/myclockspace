package server

import (
	"myclockspace.kevinmartiniano.com/internal/app"
	"myclockspace.kevinmartiniano.com/internal/routes"
)

type Routes func(*app.Router)

var register = []Routes{
	func(r *app.Router) { routes.NewHealthcheck().Register(r) },
}

func RegisterRoutes(r *app.Router) {
	for _, route := range register {
		route(r)
	}
}
