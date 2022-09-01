package v1

import (
	"errors"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ActivateCharacter(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	// Check if is in database
	var character models.Character
	if err := global.DB.First(&character, "crossbell_character_id = ?", reqCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create character
		character.CrossbellCharacterID = reqCharacterID

		// Save into database
		if err := global.DB.Create(&character).Error; err != nil {
			// Handle error
			global.Logger.Errorf("Failed to save character into database with error: %s", err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to activate character",
			})
		} else {
			// Return
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Character activated successfully",
				"result":  character,
			})
		}

	} else if err != nil {
		global.Logger.Errorf("Failed to retrieve character info from database with error: %s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to retrieve character info from database",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Character already activated",
			"result":  character,
		})
	}
}
