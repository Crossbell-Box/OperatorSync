package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	amqp "github.com/rabbitmq/amqp091-go"
)

func prepareFeedCollectQueue(queueName string, notifyClose chan *amqp.Error) (*amqp.Channel, *amqp.Queue) {

	global.Logger.Debugf("Preparing channel & queue for %s ...", queueName)

	// Prepare channel

	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Fatalf("Failed to open channel with error: %s", err.Error())
	}
	ch.NotifyClose(notifyClose) // Let caller handle reconnect

	// Prepare queue

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Fatalf("Failed to prepare queue %s with error: %s", queueName, err.Error())
	}

	return ch, &q
}
