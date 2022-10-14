package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/mq/jobs/dispatch"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func FeedCollectStartProcess() error {

	cccs := &types.ConcurrencyChannels{
		Stateful:  types.NewCtrl(config.Config.ConcurrencyStateful),
		Stateless: types.NewCtrl(config.Config.ConcurrencyStateless),
		Direct:    types.NewCtrl(config.Config.ConcurrencyDirect),
	}

	global.Logger.Debug("Feed Collect start listen on dispatching works...")

	go func() {
		for {
			// Waiting for connection
			for {
				time.Sleep(commonConsts.MQSETTINGS_ReconnectDelay)
				if commonGlobal.MQ != nil {
					break
				}
			}

			global.Logger.Infof("Preparing running environment for Feed Collect queues...")

			notifyClose := make(chan *amqp.Error)
			// Prepare channel
			ch, err := commonGlobal.MQ.Channel()
			if err != nil {
				global.Logger.Fatalf("Failed to open MQ Feeds Collect dispatch channel with error: %s", err.Error())
			}
			ch.NotifyClose(notifyClose)

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
				global.Logger.Fatalf("Failed to prepare MQ Feeds Collect dispatch queue with error: %s", err.Error())
			}

			// Prepare retrieve queue
			qRetrieve, err := ch.QueueDeclare(
				commonConsts.MQSETTINGS_FeedCollectRetrieveQueueName,
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				global.Logger.Fatalf("Failed to prepare MQ Feeds Collect retrieve queue with error: %s", err.Error())
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
				global.Logger.Fatalf("Failed to subscribe to MQ Feeds Collect dispatching queue with error: %s", err.Error())
			}
			isPendingRestart := false
			for {
				if isPendingRestart {
					break
				}
				select {
				case d := <-deliveries:
					dispatch.ProcessFeeds(cccs, ch, qRetrieve.Name, &d)
				case err := <-notifyClose:
					if err != nil {
						global.Logger.Errorf("MQ channel closed with error %d (%s), preparing to reconnect", err.Code, err.Error())
						isPendingRestart = true
					}
				}

			}
		}

	}()

	return nil
}
