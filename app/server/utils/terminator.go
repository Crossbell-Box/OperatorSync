package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
)

func TerminateAccount(account *models.Account) {
	global.Logger.Infof("Terminating account %s (%s)...", account.Username, account.Platform)
	global.DB.Delete(account)
}

func TerminateCharacter(characterId string) {
	global.Logger.Infof("Terminating character %s...", characterId)
	var accounts []models.Account
	global.DB.Find(&accounts, "crossbell_character_id = ?", characterId)
	for _, account := range accounts {
		TerminateAccount(&account)
	}
}
