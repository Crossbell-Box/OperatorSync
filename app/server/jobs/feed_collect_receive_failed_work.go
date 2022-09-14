package jobs

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func FeedCollectStartReceiveFailedWork() error {

	// Prepare channel

	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Error("Failed to open MQ Feeds Collect failed channel with error: ", err.Error())
		return err
	}

	// Prepare queue

	q, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_FeedCollectFailedQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to prepare MQ Feeds Collect failed queue with error: ", err.Error())
		return err
	}

	deliveries, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds Collect failed queue with error: ", err.Error())
		return err
	}

	global.Logger.Debug("Feed Collect start listen on failed works...")

	go func() {
		for d := range deliveries {
			feedCollectHandleFailed(&d)
		}
	}()

	return nil
}

func feedCollectHandleFailed(d *amqp.Delivery) {
	global.Logger.Warn("New failed Collect work received: ", string(d.Body))

	var workFailed commonTypes.WorkFailed
	if err := json.Unmarshal(d.Body, &workFailed); err != nil {
		global.Logger.Error("Unable to parse failed work: ", string(d.Body))
	} else {
		global.Logger.Warn("Work failed for: ", workFailed)
		global.Metrics.Work.Failed.Inc(1)
	}
}
