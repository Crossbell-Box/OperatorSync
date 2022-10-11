package v1

import (
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ActivateCharacter(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	character, err := createOrRecoverAccount(reqCharacterID)
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

func createOrRecoverAccount(crossbellCharacterID string) (*models.Character, error) {
	// Check if is in database
	var character models.Character
	if err := global.DB.Unscoped().First(&character, "crossbell_character_id = ?", crossbellCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create character
		character.CrossbellCharacterID = crossbellCharacterID

		// Save into database
		if err := global.DB.Create(&character).Error; err != nil {
			// Handle error
			global.Logger.Errorf("Failed to save character into database with error: %s", err.Error())
			return nil, fmt.Errorf("failed to activate character")
		} else {
			// Return
			return &character, nil
		}

	} else if err != nil {
		global.Logger.Errorf("Failed to retrieve character info from database with error: %s", err.Error())
		return nil, fmt.Errorf("failed to retrieve character info from database")
	} else {
		if character.DeletedAt.Valid {
			// Undelete
			global.DB.Unscoped().
				Model(&character).Where("crossbell_character_id = ?", crossbellCharacterID).
				Update("deleted_at", nil)
			// Return
			return &character, nil
		} else {
			return &character, nil
		}
	}
}
