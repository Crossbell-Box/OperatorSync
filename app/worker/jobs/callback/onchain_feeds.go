package callback

import (
	"context"
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func OnChainHandleSucceed(ch *amqp.Channel, d *amqp.Delivery, platform string, feedID uint, ipfsUri string, transaction string) {
	OnChainHandleResponse(
		ch,
		d,
		true,
		platform,
		feedID,
		ipfsUri,
		transaction,
		"",
	)
}

func OnChainHandleFailed(ch *amqp.Channel, d *amqp.Delivery, platform string, feedID uint, errMsg string) {
	OnChainHandleResponse(
		ch,
		d,
		false,
		platform,
		feedID,
		"",
		"",
		errMsg,
	)
}

func OnChainHandleResponse(ch *amqp.Channel, d *amqp.Delivery, isSucceeded bool, platform string, feedID uint, ipfsUri string, transaction string, errMsg string) {
	if isSucceeded {
		global.Logger.Debugf("New succeeded OnChain response for %s#%d: %s, %s", platform, feedID, ipfsUri, transaction)
	} else {
		global.Logger.Debugf("New failed OnChain response for %s#%d with error: %s", platform, feedID, errMsg)
	}
	response := commonTypes.OnChainRespond{
		IsSucceeded: isSucceeded,
		Platform:    platform,
		FeedID:      feedID,
		IPFSUri:     ipfsUri,
		Transaction: transaction,
	}
	if responseBytes, err := json.Marshal(&response); err != nil {
		global.Logger.Errorf("Failed to marshall OnChain respond %v with error: %s", response, err.Error())
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_PublishTimeOut)
		defer cancel()

		err = ch.PublishWithContext(
			ctx,
			"",
			d.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/json",
				CorrelationId: d.CorrelationId,
				Body:          responseBytes,
			},
		)

		if err != nil {
			global.Logger.Errorf("Failed to publish OnChain succeded respond with error: %s", err.Error())
		}
	}
	d.Ack(false)
}
