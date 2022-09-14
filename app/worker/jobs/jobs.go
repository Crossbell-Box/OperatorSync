package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/dispatch"
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
	qSucceeded, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_FeedCollectSucceededQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to prepare MQ Feeds Collect succeeded queue with error: ", err.Error())
		return err
	}

	// Prepare failed queue
	qFailed, err := ch.QueueDeclare(
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
			dispatch.ProcessFeeds(cccs, ch, qSucceeded.Name, qFailed.Name, &d)
		}
	}()

	return nil
}

func AccountValidateStartProcess() error {

	// Prepare channel
	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Error("Failed to open MQ validate channel with error: ", err.Error())
		return err
	}

	// Prepare queue
	q, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_ValidateRPCQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to declare MQ validate queue with error: ", err.Error())
		return err
	}

	// Set QoS
	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		global.Logger.Error("Failed to set validate channel QoS with error: ", err.Error())
		return err
	}

	// Register consumer
	deliveries, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ validate queue with error: ", err.Error())
		return err
	}

	go func() {
		for d := range deliveries {
			dispatch.ValidateAccounts(ch, &d)
		}
	}()

	return nil
}

func OnChainFeedsStartProcess() error {

	// Prepare channel
	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Error("Failed to open MQ Feeds OnChain dispatch channel with error: ", err.Error())
		return err
	}

	// Prepare queue
	q, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_OnChainRPCQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to declare MQ Feeds OnChain dispatch queue with error: ", err.Error())
		return err
	}

	// Set QoS
	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		global.Logger.Error("Failed to set Feeds OnChain dispatch channel QoS with error: ", err.Error())
		return err
	}

	// Register consumer
	deliveries, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ validate queue with error: ", err.Error())
		return err
	}

	go func() {
		for d := range deliveries {
			dispatch.OnChain(ch, &d)
		}
	}()

	return nil
}
