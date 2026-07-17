package server

import (
	"github.com/gin-gonic/gin"
	"myclockspace.kevinmartiniano.com/internal/routes"
)

func RegisterRoutes(routesEngine *gin.Engine) {
	routes.HealthCheckRoutes(routesEngine)
}
