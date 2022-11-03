package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/middleware"
	"github.com/gin-gonic/gin"
)

func R(e *gin.Engine) {
	e.Use(middleware.PublicCORS())

	// Public API
	publicGroup := e.Group("/")
	PublicEndpoints(publicGroup)

	// Debug Endpoints
	//if config.Config.DevelopmentMode {
	//	debugGroup := e.Group("/debug")
	//	DebugEndpoints(debugGroup)
	//}

	// API v1
	v1Group := e.Group("/v1")
	V1Endpoints(v1Group)

}
