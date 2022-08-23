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
		global.DB.Create(&character)

		// Return
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Character activated successfully",
			"result":  character,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Character already activated",
			"result":  character,
		})
	}
}
