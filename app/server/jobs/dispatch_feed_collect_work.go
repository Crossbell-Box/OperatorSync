package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func StartDispatchFeedCollectWork() {
	go func() {
		t := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-t.C:
				dispatchFeedCollectWork()
			}
		}
	}()
}

func dispatchFeedCollectWork() {
	global.Logger.Debug("Start dispatching feeds collect works...")

	nowTime := time.Now()

	// Accounts need update
	var accountsNeedUpdate []models.Account

	global.DB.Find(&accountsNeedUpdate, "next_update < ?", nowTime)

	if len(accountsNeedUpdate) == 0 {
		global.Logger.Debug("No accounts need update currently")
		return
	}

	// Dispatch update works
	for _, account := range accountsNeedUpdate {

		// Update account settings
		platform := commonConsts.SUPPORTED_PLATFORM[account.Platform]
		interv := nowTime.Sub(account.LastUpdated)
		if interv < platform.MinRefreshGap {
			interv = platform.MinRefreshGap
		} else if interv > platform.MaxRefreshGap {
			interv = platform.MaxRefreshGap
		}
		account.UpdateInterval = interv
		account.NextUpdate = nowTime.Add(interv)

		// Dispatch work
		mqChannel := commonConsts.MQSETTINGS_PlatformChannelPrefix + account.Platform
		work := commonTypes.WorkDispatched{
			DispatchAt: nowTime,
			AccountID:  account.ID,
			VerifyKey:  fmt.Sprintf("Crossbell@%s", account.CrossbellCharacter), // TODO: confirm style
			Platform:   account.Platform,
			Username:   account.Username,
			DropBefore: account.LastUpdated,
			DropAfter:  account.NextUpdate, // If cannot be peformed before DDL, work fails (cause new work would replace current one)
		}

		if workBytes, err := json.Marshal(&work); err != nil {
			global.Logger.Error("Failed to marshall work: ", work)
		} else if err = global.MQ.Publish(mqChannel, workBytes); err != nil {
			global.Logger.Error("Failed to dispatch work: ", work)
		} else {
			// Dispatched successfully
			global.Metrics.Work.Dispatched.Inc(1)
		}

		// Update account
		global.DB.Save(&account)
	}

	global.Logger.Debugf("Feeds collect works dispatched for %d accounts.", len(accountsNeedUpdate))
}
