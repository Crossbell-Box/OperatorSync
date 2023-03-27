package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func OnChainNotes(workDispatched *commonTypes.OnChainRequest, response *commonTypes.OnChainResponse) {
	global.Logger.Debug("New OnChain request received: ", workDispatched)

	// Process request
	ipfsUri, tx, characterId, noteId, err := utils.FeedOnChain(workDispatched)
	if err != nil {
		global.Logger.Error("Unable to finish OnChain request: ", workDispatched)
		OnChainHandleFailed(workDispatched.Platform, workDispatched.FeedID, err.Error(), response)
	} else {
		OnChainHandleSucceed(workDispatched.Platform, workDispatched.FeedID, ipfsUri, tx, characterId, noteId, response)
	}

}

func OnChainHandleSucceed(platform string, feedID uint, ipfsUri string, transaction string, characterId int64, noteId int64, response *commonTypes.OnChainResponse) {
	OnChainHandleResponse(
		true,
		platform,
		feedID,
		ipfsUri,
		transaction,
		characterId,
		noteId,
		"",
		response,
	)
}

func OnChainHandleFailed(platform string, feedID uint, errMsg string, response *commonTypes.OnChainResponse) {
	OnChainHandleResponse(
		false,
		platform,
		feedID,
		"",
		"",
		0,
		0,
		errMsg,
		response,
	)
}

func OnChainHandleResponse(isSucceeded bool, platform string, feedID uint, ipfsUri string, transaction string, characterId int64, noteId int64, errMsg string, response *commonTypes.OnChainResponse) {
	*response = commonTypes.OnChainResponse{
		IsSucceeded: isSucceeded,
		Platform:    platform,
		FeedID:      feedID,
		IPFSUri:     ipfsUri,
		Transaction: transaction,
		CharacterID: characterId,
		NoteID:      noteId,
		Message:     errMsg,
	}
}
