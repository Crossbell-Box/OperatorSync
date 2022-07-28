package v1

import (
	"errors"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CheckCharacter(ctx *gin.Context) {
	// Parse request params
	reqCharacter := ctx.Param("character")

	// Check if is in database
	var character models.Character
	if err := global.DB.First(&character, "crossbell_character = ?", reqCharacter).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Character not enabled",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Character found",
			"result":  character,
		})
	}

}
