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

	v1Group := e.Group("/v1")
	V1Endpoints(v1Group)

}
