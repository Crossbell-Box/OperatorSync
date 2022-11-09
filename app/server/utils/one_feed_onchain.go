package utils

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func OneFeedOnChain(account *models.Account, feed *models.Feed) (string, string, error, bool) {

	onChainRequest := types.OnChainRequest{
		FeedID:               feed.ID,
		CrossbellCharacterID: account.CrossbellCharacterID,
		Platform:             account.Platform,
		Username:             account.Username,
		RawFeed:              feed.RawFeed,
	}

	var onChainResponse types.OnChainResponse

	// Start request
	errChan := make(chan error, 1)

	go func() {
		errChan <- global.RPC.Call(
			fmt.Sprintf("%s.%s", consts.RPCSETTINGS_BaseServiceName, consts.RPCSETTINGS_OnChainServiceName),
			onChainRequest,
			&onChainResponse,
		)
	}()

	// Set timeout
	select {
	case <-time.After(consts.RPCSETTINGS_OnChainRequestTimeOut):
		// Timeout
		global.Logger.Errorf("OnChain request timeout...")
		return "", "", fmt.Errorf("onChain request timeout"), false

	case err := <-errChan:
		if err != nil {
			global.Logger.Errorf("Failed to receive on chain respond with error: %s", err.Error())
			AccountOnChainPause(account, fmt.Sprintf("Failed toreceive on chain respond with error: %s", err.Error()))
			return "", "", err, false
		}
	}

	// Validate response
	if !onChainResponse.IsSucceeded {
		global.Logger.Errorf("Failed to finish OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, onChainResponse.Message)
		AccountOnChainPause(account, fmt.Sprintf("Failed to finish OnChain work for feed %s#%d", account.Platform, feed.ID))
		isAccountTerminated := CheckAndTerminateIfNeed(account)
		return onChainResponse.IPFSUri, onChainResponse.Transaction, fmt.Errorf(onChainResponse.Message), isAccountTerminated
	}

	return onChainResponse.IPFSUri, onChainResponse.Transaction, nil, false

}
