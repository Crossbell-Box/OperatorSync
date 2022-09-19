package jobs

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func feedOnChainDispatchWork(account *models.Account, feeds []models.Feed) {
	if utils.IsAccountOnChainPaused(account) {
		global.Logger.Errorf("OnChain process has been paused for account %s#%d with reason: %s", account.Platform, account.ID, account.OnChainPauseMessage)
		return
	}

	for index, feed := range feeds {

		ipfsUri, tx, err := OneFeedOnChain(account, &feed)
		if err != nil {
			global.Logger.Errorf("Failed to push all feeds on chain with error: %s", err.Error())
			break
		} else {
			// `feed` is read-only and cannot change its value
			feeds[index].IPFSUri = ipfsUri
			feeds[index].Transaction = tx
			account.NotesCount++
		}

	}

	// Update feed save into database
	if err := global.DB.Scopes(models.FeedTable(models.Feed{
		Feed: types.Feed{
			Platform: account.Platform,
		},
	})).Save(&feeds).Error; err != nil {
		global.Logger.Error("Failed to save updated feeds", feeds)
	}

	// Update account
	global.DB.Save(account)
}

func OneFeedOnChain(account *models.Account, feed *models.Feed) (string, string, error) {

	onChainRequest := commonTypes.OnChainRequest{
		FeedID:               feed.ID,
		CrossbellCharacterID: account.CrossbellCharacterID,
		Platform:             account.Platform,
		Username:             account.Username,
		RawFeed:              feed.RawFeed,
	}

	var onChainResponse commonTypes.OnChainResponse

	// Start request
	errChan := make(chan error, 1)

	go func() {
		errChan <- global.RPC.Call(
			fmt.Sprintf("%s.%s", commonConsts.RPCSETTINGS_BaseServiceName, commonConsts.RPCSETTINGS_OnChainServiceName),
			onChainRequest,
			&onChainResponse,
		)
	}()

	// Set timeout
	select {
	case <-time.After(commonConsts.RPCSETTINGS_OnChainRequestTimeOut):
		// Timeout
		global.Logger.Errorf("OnChain request timeout...")
		return "", "", fmt.Errorf("onChain request timeout")

	case err := <-errChan:
		if err != nil {
			global.Logger.Errorf("Failed to receive on chain respond with error: %s", err.Error())
			utils.AccountOnChainPause(account, fmt.Sprintf("Failed toreceive on chain respond with error: %s", err.Error()))
			return "", "", err
		}
	}

	// Validate response
	if !onChainResponse.IsSucceeded {
		global.Logger.Errorf("Failed to finish OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, onChainResponse.Message)
		utils.AccountOnChainPause(account, fmt.Sprintf("Failed to finish OnChain work for feed %s#%d", account.Platform, feed.ID))
		return onChainResponse.IPFSUri, onChainResponse.Transaction, fmt.Errorf(onChainResponse.Message)
	}

	return onChainResponse.IPFSUri, onChainResponse.Transaction, nil

}
