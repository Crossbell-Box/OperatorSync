package routers

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/handlers/v1"
	"github.com/gin-gonic/gin"
)

func V1Endpoints(rg *gin.RouterGroup) {
	rg.GET("/:character", v1.CheckCharacter)
	rg.POST("/:character", v1.ActivateCharacter)
	rg.DELETE("/:character", v1.DeactivateAccount)
	rg.GET("/:character/account", v1.ListAccounts)
	rg.POST("/:character/account/bind/:platform/:username", v1.BindAccount)
	rg.POST("/:character/account/sync/:platform/:username", v1.ForceSyncAccount)
	rg.DELETE("/:character/account/unbind/:platform/:username", v1.UnbindAccount)
	rg.GET("/:character/media", v1.ListMedias)

	rg.GET("/feed/:platform/:username", v1.ListSingleAccountFeeds)
}
