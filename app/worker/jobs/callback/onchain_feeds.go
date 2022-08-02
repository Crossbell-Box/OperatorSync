package callback

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func OnChainHandleSucceed(mqReply string, platform string, feedID uint, ipfsUri string, transaction string) {
	// Only succeeded is allowed, or please retry
	global.Logger.Debugf("New succeeded OnChain response for %s#%d: %s, %s", platform, feedID, ipfsUri, transaction)
	if workBytes, err := json.Marshal(&commonTypes.OnChainRespond{
		IsSucceeded: true,
		Platform:    platform,
		FeedID:      feedID,
		IPFSUri:     ipfsUri,
		Transaction: transaction,
	}); err != nil {
		global.Logger.Errorf("Failed to parse OnChain respond for feed %s#%d with error: %s", platform, feedID, err.Error())
	} else {
		// Dispatch with ID
		if err := global.MQ.Publish(mqReply, workBytes); err != nil {
			global.Logger.Errorf("Failed to publish OnChain succeded respond with error: %s", err.Error())
		}
	}
}

func OnChainHandleFailed(mqReply string, platform string, feedID uint, errMsg string) {
	// Only succeeded is allowed, or please retry
	global.Logger.Debugf("New failed OnChain response for %s#%d with error: %s", platform, feedID, errMsg)
	if workBytes, err := json.Marshal(&commonTypes.OnChainRespond{
		IsSucceeded: false,
		Platform:    platform,
		FeedID:      feedID,
		Message:     errMsg,
	}); err != nil {
		global.Logger.Errorf("Failed to parse OnChain respond for feed %s#%d with error: %s", platform, feedID, err.Error())
	} else {
		// Dispatch with ID
		if err := global.MQ.Publish(mqReply, workBytes); err != nil {
			global.Logger.Errorf("Failed to publish OnChain succeded respond with error: %s", err.Error())
		}
	}
}
