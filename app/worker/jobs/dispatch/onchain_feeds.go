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
	global.Logger.Debug("New OnChain Dispatched received: ", string(m.Data))

	// Report received progress
	if err := m.InProgress(); err != nil {
		global.Logger.Error("Failed to report work progress")
	}

	// Parse request
	var workDispatched commonTypes.OnChainDispatched
	if err := json.Unmarshal(m.Data, &workDispatched); err != nil {
		global.Logger.Error("Unable to parse dispatched: ", string(m.Data))
		if err := m.Nak(); err != nil {
			global.Logger.Error("Failed to report work Nak status")
		}
		return
	}

	// Process request
	ipfsUri, tx, err := utils.FeedOnChain(&workDispatched)
	if err != nil {
		global.Logger.Error("Unable to finish onchain request: ", string(m.Data))
		if err := m.Nak(); err != nil {
			global.Logger.Error("Failed to report work Nak status")
		}
		return
	}

	// Report uploaded progress
	if err := m.InProgress(); err != nil {
		global.Logger.Error("Failed to report work progress")
	}

	callback.OnChainHandleSucceed(workDispatched.Platform, workDispatched.FeedID, ipfsUri, tx)

	// Report finish progress
	if err := m.Ack(); err != nil {
		global.Logger.Error("Failed to report work Ack status")
	}
}
