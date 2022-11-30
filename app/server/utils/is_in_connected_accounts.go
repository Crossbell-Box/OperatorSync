package utils

import (
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func IsInConnectedAccounts(platform string, username string, connectedAccounts []commonTypes.Account) bool {
	for _, account := range connectedAccounts {
		if account.Platform == platform && account.Username == username {
			return true
		}
	}
	// No match
	return false
}
