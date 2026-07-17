package routes

import (
	"github.com/gin-gonic/gin"
	"myclockspace.kevinmartiniano.com/internal/healthcheck"
)

func HealthCheckRoutes(routes *gin.Engine) {
	routes.GET("/healthcheck", healthcheck.Index)
}
