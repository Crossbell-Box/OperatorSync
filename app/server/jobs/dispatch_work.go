package jobs

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strings"
	"time"
)

func StartDispatchFeedCollectWork() {
	go func() {
		t := time.NewTimer(1 * time.Minute)
		for {
			select {
			case <-t.C:
				dispatch()
			}
		}
	}()
}

func dispatch() {
	global.Logger.Debug("Start dispatching feeds collect works...")

	nowTime := time.Now()

	// Accounts need update
	var accountsNeedUpdate []models.Account
	global.DB.Find(&accountsNeedUpdate, "next_update < ?", nowTime)

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
			DispatchAt:  nowTime,
			AccountID:   account.ID,
			Platform:    account.Platform,
			CollectLink: strings.ReplaceAll(platform.FeedLink, "{{username}}", account.Username),
			DropBefore:  account.LastUpdated,
			DropAfter:   account.NextUpdate, // If cannot be peformed before DDL, work fails (cause new work would replace current one)
		}

		if workBytes, err := json.Marshal(&work); err != nil {
			global.Logger.Error("Failed to marshall work: ", work)
		} else if err = global.MQ.Publish(mqChannel, workBytes); err != nil {
			global.Logger.Error("Failed to dispatch work: ", work)
		} else {
			// Dispatched successfully
			global.MetricsDispatchedWorkCount.Inc(1)
		}
	}

	// Update accounts
	global.DB.Save(&accountsNeedUpdate)

	global.Logger.Debug("Feeds collect works dispatched!")
}
