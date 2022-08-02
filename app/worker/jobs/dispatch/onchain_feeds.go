package dispatch

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
)

func OnChain(m *nats.Msg) {
	global.Logger.Debug("New OnChain request received: ", string(m.Data))

	// Parse request
	var workDispatched commonTypes.OnChainRequest
	if err := json.Unmarshal(m.Data, &workDispatched); err != nil {
		global.Logger.Error("Unable to parse request: ", string(m.Data))
		// Even unable to callback
		return
	}

	// Process request
	ipfsUri, tx, err := utils.FeedOnChain(&workDispatched)
	if err != nil {
		global.Logger.Error("Unable to finish onchain request: ", string(m.Data))
		callback.OnChainHandleFailed(m.Reply, workDispatched.Platform, workDispatched.FeedID, err.Error())
		return
	}

	callback.OnChainHandleSucceed(m.Reply, workDispatched.Platform, workDispatched.FeedID, ipfsUri, tx)
}
