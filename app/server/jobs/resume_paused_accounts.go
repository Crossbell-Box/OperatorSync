package jobs

import (
	"errors"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	"gorm.io/gorm"
	"net/http"
	"sort"
	"time"
)

func ResumePausedAccounts() {
	global.Logger.Debug("Paused account resume work start dispatching...")
	go func() {
		t := time.NewTicker(consts.JOBS_INTERVAL_RESUME_PAUSED_ACCOUNTS)
		for {
			select {
			case <-t.C:
				TryToResumeAllPausedAccounts()
			}
		}
	}()
}

var (
	_isResumeWorkProcessing bool
)

func init() {
	_isResumeWorkProcessing = false
}

func TryToResumeAllPausedAccounts() {
	if _isResumeWorkProcessing {
		// No need to start another one, skip
		return
	}

	// Set busy flag
	_isResumeWorkProcessing = true

	global.Logger.Debug("Start trying to resume all paused accounts...")

	config.Status.Jobs.ResumePausedAccountsLastRun = time.Now()

	if config.Config.HeartBeatWebhooks.AccountResume != "" {
		// Send heartbeat packet
		global.Logger.Debug("Sending account resume heartbeat packet")
		_, err := (&http.Client{}).Get(config.Config.HeartBeatWebhooks.AccountResume) // Ignore response
		if err != nil {
			global.Logger.Errorf("Failed to send account resume heartbeat packet with error: %s", err.Error())
		}
	}

	var pausedAccounts []models.Account

	global.DB.Find(&pausedAccounts, "is_onchain_paused = ?", true)

	for _, pa := range pausedAccounts {
		if tryToResumeOnePausedAccount(&pa) {
			// Recovered
			global.Logger.Debugf("Account #%d (%s@%s) recovered", pa.ID, pa.Username, pa.Platform)
			utils.AccountOnChainResume(&pa)
		} else {
			global.Logger.Errorf("Still unable to resume account #%d (%s@%s)", pa.ID, pa.Username, pa.Platform)

		}
	}

	// Unset busy flag
	_isResumeWorkProcessing = false

	global.Logger.Debug("All paused accounts checked, and already tried best to resume them.")
}

func tryToResumeOnePausedAccount(account *models.Account) bool {
	// Try to find the earliest paused feed
	global.Logger.Debugf("Try to find the first failed feed of account %s#%d", account.Platform, account.ID)

	var pausedFeeds models.FeedsArray
	if err := global.DB.Scopes(models.FeedTable(models.Feed{
		Feed: types.Feed{
			Platform: account.Platform,
		},
	})).Where("account_id = ?", account.ID).Find(&pausedFeeds, "transaction = ? OR transaction IS NULL", "").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Something is wrong
			global.Logger.Errorf("Something is wrong, failed to find the first paused feed of account %s#%d", account.Platform, account.ID)
			// But should be resumable, maybe?
			return true
		} else {
			global.Logger.Errorf("Failed to get first failed feed from database with error: %s", err.Error())
			return false
		}
	}

	sort.Sort(pausedFeeds)

	isAllSucceeded := true
	isAccountTerminated := false

	for index, feed := range pausedFeeds {
		// Recover feeds' media
		for _, mediaIPFSUri := range feed.MediaIPFSUris {
			var media models.Media
			global.DB.First(&media, "ipfs_uri = ?", mediaIPFSUri)
			feed.Media = append(feed.Media, media.Media)
		}

		// Try to push as many feeds as we can
		ipfsUri, tx, err, terminated := utils.OneFeedOnChain(account, &feed)
		if err != nil {
			// Oops
			global.Logger.Errorf("Failed to OnCHain feed %s#%d with error: %s", account.Platform, feed.ID, err.Error())
			isAllSucceeded = false
			if terminated {
				isAccountTerminated = true
			}
			break
		} else {
			global.Logger.Debugf("Succeeded to OnChain feed %s#%d", account.Platform, feed.ID)
			pausedFeeds[index].IPFSUri = ipfsUri
			pausedFeeds[index].Transaction = tx

			// Increase succeeded notes count
			account.NotesCount++
		}
	}

	// Save DB
	global.DB.Scopes(models.FeedTable(models.Feed{
		Feed: types.Feed{
			Platform: account.Platform,
		},
	})).Save(&pausedFeeds)

	if !isAccountTerminated {
		// Update account if not terminated
		global.DB.Save(account)
	}

	return isAllSucceeded

}
