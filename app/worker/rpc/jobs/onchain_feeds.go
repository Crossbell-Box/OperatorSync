package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func OnChainNotes(workDispatched *commonTypes.OnChainRequest) *commonTypes.OnChainResponse {
	global.Logger.Debug("New OnChain request received: ", workDispatched)

	// Process request
	ipfsUri, tx, err := utils.FeedOnChain(workDispatched)
	if err != nil {
		global.Logger.Error("Unable to finish onChain request: ", workDispatched)
		return OnChainHandleFailed(workDispatched.Platform, workDispatched.FeedID, err.Error())
	}

	return OnChainHandleSucceed(workDispatched.Platform, workDispatched.FeedID, ipfsUri, tx)
}

func OnChainHandleSucceed(platform string, feedID uint, ipfsUri string, transaction string) *commonTypes.OnChainResponse {
	return OnChainHandleResponse(
		true,
		platform,
		feedID,
		ipfsUri,
		transaction,
		"",
	)
}

func OnChainHandleFailed(platform string, feedID uint, errMsg string) *commonTypes.OnChainResponse {
	return OnChainHandleResponse(
		false,
		platform,
		feedID,
		"",
		"",
		errMsg,
	)
}

func OnChainHandleResponse(isSucceeded bool, platform string, feedID uint, ipfsUri string, transaction string, errMsg string) *commonTypes.OnChainResponse {
	return &commonTypes.OnChainResponse{
		IsSucceeded: isSucceeded,
		Platform:    platform,
		FeedID:      feedID,
		IPFSUri:     ipfsUri,
		Transaction: transaction,
		Message:     errMsg,
	}
}
