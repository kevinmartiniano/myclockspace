package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	ginDefault := gin.Default()

	RegisterRoutes(ginDefault)

	ginDefault.Run(":8080")
}
