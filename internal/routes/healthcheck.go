package routes

import (
	"net/http"

	"myclockspace.kevinmartiniano.com/internal/healthcheck"
)

func HealthCheckRoutes() {
	http.HandleFunc("/healthcheck", healthcheck.Index)
}
