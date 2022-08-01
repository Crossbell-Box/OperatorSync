package callback

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
)

func OnChainHandleSucceed(platform string, feedID uint, ipfsUri string, transaction string) {
	// Only succeeded is allowed, or please retry
	global.Logger.Debugf("New succeeded OnChain response for %s#%d: %s, %s", platform, feedID, ipfsUri, transaction)
	if workBytes, err := json.Marshal(&commonTypes.OnChainRespond{
		Platform:    platform,
		FeedID:      feedID,
		IPFSUri:     ipfsUri,
		Transaction: transaction,
	}); err != nil {
		global.Logger.Errorf("Failed to parse OnChain respond for feed %s#%d with error: %s", platform, feedID, err.Error())
	} else {
		// Dispatch with ID
		if _, err := global.MQJS.Publish(
			fmt.Sprintf("%s.%s", commonConsts.MQSETTINGS_OnChainStreamName, commonConsts.MQSETTINGS_OnChainResponseSubjectName),
			workBytes,
			nats.MsgId(fmt.Sprintf("%s#%d-resp", platform, feedID)),
		); err != nil {
			global.Logger.Errorf("Failed to publish OnChain succeded respond with error: %s", err.Error())
		}
	}
}
