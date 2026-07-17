package server

import (
	"myclockspace.kevinmartiniano.com/internal/routes"
)

func RegisterRoutes() {
	routes.HealthCheckRoutes()
}
