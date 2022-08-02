package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
)

func feedOnChainDispatchWork(account *models.Account, feeds []models.Feed) {
	for _, feed := range feeds {
		if workBytes, err := json.Marshal(&commonTypes.OnChainDispatched{
			FeedID:               feed.ID,
			CrossbellCharacterID: account.CrossbellCharacterID,
			Platform:             feed.Platform,
			Username:             account.Username,
			RawFeed:              feed.RawFeed,
		}); err != nil {
			global.Logger.Errorf("Failed to parse OnChain work for feed %s#%d with error: %s", feed.Platform, feed.ID, err.Error())
		} else {
			// Dispatch with ID
			if _, err := global.MQJS.Publish(
				fmt.Sprintf("%s.%s", commonConsts.MQSETTINGS_OnChainStreamName, commonConsts.MQSETTINGS_OnChainDispatchSubjectName),
				workBytes,
				nats.MsgId(fmt.Sprintf("%s#%d-disp", feed.Platform, feed.ID)),
			); err != nil {
				global.Logger.Errorf("Failed to dispatch OnChain work for feed %s#%d with error: %s", feed.Platform, feed.ID, err.Error())
			}
		}
	}
}
