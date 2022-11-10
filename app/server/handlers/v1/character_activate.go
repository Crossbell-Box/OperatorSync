package v1

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActivateCharacter(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	character, err := utils.CreateOrRecoverCharacter(reqCharacterID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": fmt.Sprintf("Failed to activate character with error: %s", err.Error()),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Character activated successfully",
			"result":  character,
		})
	}
}
