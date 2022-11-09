package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
)

func CheckAndTerminateIfNeed(account *models.Account) bool {
	// Check operator for character
	if isOperatorValid, checkOperatorErr := CheckOperator(account.CrossbellCharacterID); checkOperatorErr != nil {
		// Failed to check operator status, unable to handle now
		return false
	} else if !isOperatorValid {
		// Operator is no longer valid for this character
		TerminateCharacter(account.CrossbellCharacterID)
		return true
	} else if isAccountValid, validateAccountErr := ValidateAccount(account.CrossbellCharacterID, account.Platform, account.Username); validateAccountErr != nil {
		// Operator is valid, but failed to check account validate string status, unable to handle now
		return false
	} else if !isAccountValid {
		// Account validate string has been removed
		TerminateAccount(account)
		return true
	} else {
		// It's just fine
		return false
	}
}
