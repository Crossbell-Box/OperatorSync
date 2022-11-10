package v1

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func BindAccount(ctx *gin.Context) {
	// Only accept when pass validate

	reqCharacterID := ctx.Param("character") // ID
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	// User selectable start time
	nowTimeStr := time.Unix(0, 0).Format(time.RFC3339)
	startFromStr := ctx.DefaultQuery("from", nowTimeStr)
	startFromTime, err := time.Parse(time.RFC3339, startFromStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": fmt.Sprintf("Failed to parse time with error (%s), please use RFC3339 format (%s)", err.Error(), nowTimeStr),
		})
		return
	}
	isAccountInherit := strings.Contains(strings.ToLower(ctx.DefaultQuery("inherit", "true")), "t")

	props := utils.BindAccountProps{
		CrossbellCharacterID: reqCharacterID,
		Platform:             reqPlatform,
		Username:             reqUsername,
		StartFrom:            startFromTime,
		IsInherit:            isAccountInherit,
	}

	result, message, err := utils.AccountBind(&props)
	if err != nil {
		global.Logger.Errorf("Failed to bind account %s (%s) for character #%s with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
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
