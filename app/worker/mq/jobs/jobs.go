package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/mq/jobs/dispatch"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
)

func FeedCollectStartProcess() error {

	cccs := &types.ConcurrencyChannels{
		Stateful:  types.NewCtrl(config.Config.ConcurrencyStateful),
		Stateless: types.NewCtrl(config.Config.ConcurrencyStateless),
		Direct:    types.NewCtrl(config.Config.ConcurrencyDirect),
	}

	// Prepare channel
	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Error("Failed to open MQ Feeds Collect dispatch channel with error: ", err.Error())
		return err
	}

	// Prepare dispatch queue
	qDispatch, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_FeedCollectDispatchQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to prepare MQ Feeds Collect dispatch queue with error: ", err.Error())
		return err
	}

	// Prepare succeeded queue
	qRetrieve, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_FeedCollectRetrieveQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to prepare MQ Feeds Collect retrieve queue with error: ", err.Error())
		return err
	}

	deliveries, err := ch.Consume(
		qDispatch.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds Collect dispatching queue with error: ", err.Error())
		return err
	}

	global.Logger.Debug("Feed Collect start listen on dispatching works...")

	go func() {
		for d := range deliveries {
			dispatch.ProcessFeeds(cccs, ch, qRetrieve.Name, &d)
		}
	}()

	return nil
}
