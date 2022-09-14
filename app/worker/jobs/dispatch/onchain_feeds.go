package dispatch

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func OnChain(ch *amqp.Channel, d *amqp.Delivery) {
	global.Logger.Debug("New OnChain request received: ", string(d.Body))

	// Parse request
	var workDispatched commonTypes.OnChainRequest
	if err := json.Unmarshal(d.Body, &workDispatched); err != nil {
		global.Logger.Error("Unable to parse request: ", string(d.Body))
		// Even unable to callback
		return
	}

	// Process request
	ipfsUri, tx, err := utils.FeedOnChain(&workDispatched)
	if err != nil {
		global.Logger.Error("Unable to finish onchain request: ", string(d.Body))
		callback.OnChainHandleFailed(ch, d, workDispatched.Platform, workDispatched.FeedID, err.Error())
		return
	}

	callback.OnChainHandleSucceed(ch, d, workDispatched.Platform, workDispatched.FeedID, ipfsUri, tx)
}
