package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/handlers/public"
	"github.com/gin-gonic/gin"
)

func PublicEndpoints(rg *gin.RouterGroup) {
	// Health Check
	rg.GET("/healthcheck", public.HealthCheck)

	// List
	rg.GET("/account/:character", public.ListAccounts)
	rg.POST("/account/:character/refresh", public.RefreshAccounts)
	rg.GET("/feed/:platform/:username", public.ListFeeds)
}
