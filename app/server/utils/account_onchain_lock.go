package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"time"
)

func IsAccountOnChainPaused(account *models.Account) bool {
	return account.IsOnChainPaused
}

func AccountOnChainPause(account *models.Account, message string) {
	account.IsOnChainPaused = true
	account.OnChainPausedAt = time.Now()
	account.OnChainPauseMessage = message

	global.DB.Save(&account)
}

func AccountOnChainResume(account *models.Account) {
	account.IsOnChainPaused = false

	global.DB.Save(&account)
}
