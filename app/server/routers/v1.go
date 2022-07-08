package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/handlers/v1"
	"github.com/gin-gonic/gin"
)

func V1Endpoints(rg *gin.RouterGroup) {
	rg.GET("/account/:character", v1.ListAccounts)
	rg.POST("/account/:character/refresh", v1.RefreshAccounts)
	rg.GET("/feed/:platform/:username", v1.ListFeeds)
	rg.GET("/media/:character", v1.ListMedias)
}
