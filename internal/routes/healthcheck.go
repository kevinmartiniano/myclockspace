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
		Register: HealthCheckRoutes,
	}
}

func HealthCheckRoutes(router *app.Router) {
	router.Get("/healthcheck", healthcheck.Index)
}
