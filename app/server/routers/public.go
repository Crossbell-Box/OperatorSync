package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/handlers/public"
	"github.com/gin-gonic/gin"
)

func PublicEndpoints(rg *gin.RouterGroup) {
	// List endpoints
	rg.GET("/", public.ListEndpoints)

	// Health Check
	rg.GET("/healthcheck", public.HealthCheck)

	// List platforms
	rg.GET("/platforms", public.ListPlatforms)

	// Metrics
	rg.GET("/metrics", public.Metrics)

}
