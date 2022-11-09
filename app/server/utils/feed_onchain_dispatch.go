package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
)

func FeedOnChainDispatchWork(account *models.Account, feeds []models.Feed) {
	if IsAccountOnChainPaused(account) {
		global.Logger.Errorf("OnChain process has been paused for account %s#%d with reason: %s", account.Platform, account.ID, account.OnChainPauseMessage)
		return
	}

	isAccountTerminated := false

	for index, feed := range feeds {

		ipfsUri, tx, err, terminated := OneFeedOnChain(account, &feed)
		if err != nil {
			global.Logger.Errorf("Failed to push all feeds on chain with error: %s", err.Error())
			if terminated {
				isAccountTerminated = true
			}
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

	if !isAccountTerminated {
		// Update account if not terminated
		global.DB.Save(account)
	}
}
