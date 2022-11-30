package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/chain"
	"github.com/Crossbell-Box/OperatorSync/app/worker/indexer"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strings"
)

func CheckOnChainData(workDispatched *commonTypes.CheckOnChainDataRequest, response *commonTypes.CheckOnChainDataResponse) {
	// Check if operator is set
	isOperatorSet, err := chain.CheckOperator(workDispatched.CrossbellCharacterID)
	if err != nil {
		*response = commonTypes.CheckOnChainDataResponse{
			IsSucceeded:     false,
			Message:         err.Error(),
			IsOperatorValid: false,
			//IsAccountConnected:   false,
		}
		return
	} else if !isOperatorSet {
		*response = commonTypes.CheckOnChainDataResponse{
			IsSucceeded:     true,
			Message:         "Operator not set",
			IsOperatorValid: false,
			//IsAccountConnected:   false,
		}
		return
	}

	// Check if platform is set in metadata
	metadata, err := indexer.GetCharacterMetadataFromIndexer(workDispatched.CrossbellCharacterID)
	if err != nil {
		*response = commonTypes.CheckOnChainDataResponse{
			IsSucceeded:       false,
			Message:           err.Error(),
			IsOperatorValid:   true,
			ConnectedAccounts: nil,
		}
		return
	}

	*response = commonTypes.CheckOnChainDataResponse{
		IsSucceeded:       true,
		Message:           "",
		IsOperatorValid:   true,
		ConnectedAccounts: nil,
	}

	for _, connectedAccount := range metadata.Content.ConnectedAccounts {
		// Parse accounts
		accountPlatformSplits := strings.Split(strings.Replace(connectedAccount, "csb://account:", "", 1), "@")
		response.ConnectedAccounts = append(response.ConnectedAccounts, commonTypes.Account{
			Platform: accountPlatformSplits[len(accountPlatformSplits)-1],
			Username: strings.Join(accountPlatformSplits[0:len(accountPlatformSplits)-1], "@"), // For fediverse
		})
	}

}
