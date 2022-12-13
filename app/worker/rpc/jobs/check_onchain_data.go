package jobs

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/chain"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/indexer"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
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
		connectedAccountStr := ""
		if strVal, isString := connectedAccount.(string); isString {
			// Is string
			connectedAccountStr = strVal
		} else {
			// Check if is ConnectedAccountsObject
			connectedAccountBytes, err := json.Marshal(&connectedAccount)
			if err != nil {
				// Failed
				global.Logger.Errorf("Failed to marshal connected account %v with error: %v", connectedAccount, err)
				continue
			}
			var connectedAccountObject types.ConnectedAccountsObject
			err = json.Unmarshal(connectedAccountBytes, &connectedAccountObject)
			if err != nil {
				// Not object
				global.Logger.Errorf("Failed to unmarshall connected account %v with error: %v", connectedAccount, err)
				continue
			}

			// Finally
			connectedAccountStr = connectedAccountObject.URI
		}

		// Not empty
		if connectedAccountStr != "" {
			global.Logger.Debugf("New connected account found: %s", connectedAccountStr)

			accountPlatformSplits := strings.Split(strings.Replace(connectedAccountStr, "csb://account:", "", 1), "@")
			response.ConnectedAccounts = append(response.ConnectedAccounts, commonTypes.Account{
				Platform: accountPlatformSplits[len(accountPlatformSplits)-1],
				Username: strings.Join(accountPlatformSplits[0:len(accountPlatformSplits)-1], "@"), // For fediverse
			})
		}
	}

}
