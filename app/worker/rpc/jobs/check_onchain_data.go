package jobs

import (
	"fmt"
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
			IsSucceeded:        false,
			Message:            err.Error(),
			IsOperatorValid:    true,
			IsAccountConnected: false,
		}
		return
	}

	targetConnectedAccount := strings.ToLower(fmt.Sprintf("csb://account:%s@%s", workDispatched.Account, workDispatched.Platform))
	for _, connectedAccount := range metadata.Content.ConnectedAccounts {
		if strings.ToLower(connectedAccount) == targetConnectedAccount {
			*response = commonTypes.CheckOnChainDataResponse{
				IsSucceeded:        true,
				Message:            "",
				IsOperatorValid:    true,
				IsAccountConnected: true,
			}
			return
		}
	}

	// Else
	*response = commonTypes.CheckOnChainDataResponse{
		IsSucceeded:        true,
		Message:            "Connected account not set",
		IsOperatorValid:    true,
		IsAccountConnected: false,
	}
}
