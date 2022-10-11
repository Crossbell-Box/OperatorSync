package v1

import (
	"errors"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func DeactivateAccount(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	// Check if is in database
	var character models.Character
	if err := global.DB.First(&character, "crossbell_character_id = ?", reqCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create character
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Character not activated",
			"result":  true,
		})

	} else if err != nil {
		global.Logger.Errorf("Failed to retrieve character info from database with error: %s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to retrieve character info from database",
		})
	} else {
		// Check if any active sync account

		if err := global.DB.First(&models.Account{}, "crossbell_character_id = ?", reqCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			// Delete from database
			if err := global.DB.Delete(&character, "crossbell_character_id = ?", reqCharacterID).Error; err != nil {
				// Handle error
				global.Logger.Errorf("Failed to delete character from database with error: %s", err.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"ok":      false,
					"message": "Failed to deactivate character",
				})
			} else {
				// Return
				ctx.JSON(http.StatusOK, gin.H{
					"ok":      true,
					"message": "Character deactivated successfully",
					"result":  true,
				})
			}
		} else if err != nil {
			// Handle error
			global.Logger.Errorf("Failed to check character's accounts status from database with error: %s", err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to check character's accounts status",
			})
		} else {
			// Not clear
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Character still have active sync accounts bind",
				"result":  false,
			})
		}
	}

}
