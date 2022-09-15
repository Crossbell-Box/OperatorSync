package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func OnChainNotes(workDispatched *commonTypes.OnChainRequest) *commonTypes.OnChainRespond {
	global.Logger.Debug("New OnChain request received: ", workDispatched)

	// Process request
	ipfsUri, tx, err := utils.FeedOnChain(workDispatched)
	if err != nil {
		global.Logger.Error("Unable to finish onChain request: ", workDispatched)
		return OnChainHandleFailed(workDispatched.Platform, workDispatched.FeedID, err.Error())
	}

	return OnChainHandleSucceed(workDispatched.Platform, workDispatched.FeedID, ipfsUri, tx)
}

func OnChainHandleSucceed(platform string, feedID uint, ipfsUri string, transaction string) *commonTypes.OnChainRespond {
	return OnChainHandleResponse(
		true,
		platform,
		feedID,
		ipfsUri,
		transaction,
		"",
	)
}

func OnChainHandleFailed(platform string, feedID uint, errMsg string) *commonTypes.OnChainRespond {
	return OnChainHandleResponse(
		false,
		platform,
		feedID,
		"",
		"",
		errMsg,
	)
}

func OnChainHandleResponse(isSucceeded bool, platform string, feedID uint, ipfsUri string, transaction string, errMsg string) *commonTypes.OnChainRespond {
	return &commonTypes.OnChainRespond{
		IsSucceeded: isSucceeded,
		Platform:    platform,
		FeedID:      feedID,
		IPFSUri:     ipfsUri,
		Transaction: transaction,
		Message:     errMsg,
	}
}
