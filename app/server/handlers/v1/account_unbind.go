package v1

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnbindAccount(ctx *gin.Context) {
	// Only accept when fails validate

	reqCharacterID := ctx.Param("character")
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	if _, ok := commonConsts.SUPPORTED_PLATFORM[reqPlatform]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Platform not supported",
		})
		return
	}

	props := utils.UnbindAccountProps{
		CrossbellCharacterID: reqCharacterID,
		Platform:             reqPlatform,
		Username:             reqUsername,
	}

	result, message, err := utils.AccountUnbind(&props)
	if err != nil {
		global.Logger.Errorf("Failed to unbind account %s (%s) for character #%s with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": message,
			"result":  result,
		})
	}

}
