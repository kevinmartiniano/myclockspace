package routes

import (
	"myclockspace.kevinmartiniano.com/internal/app"
	"myclockspace.kevinmartiniano.com/internal/healthcheck"
)

type Healthcheck struct {
	Register func(*app.Router)
}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{
		Register: func(router *app.Router) { routes(router) },
	}
}

func routes(router *app.Router) {
	handler := healthcheck.NewHealthcheckHandler()
	
	router.Get("/healthcheck", handler.Index)
}
