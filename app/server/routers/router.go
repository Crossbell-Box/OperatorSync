package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/middleware"
	"github.com/gin-gonic/gin"
)

func R(e *gin.Engine) {
	e.Use(middleware.PublicCORS())

	// Public API
	publicEndpoint := e.Group("/")
	PublicEndpoints(publicEndpoint)

}
