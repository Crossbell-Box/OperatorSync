package jobs

import (
	"context"
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func FeedCollectStartDispatchWork() {

	// Prepare channel

	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Fatal("Failed to open MQ Feeds Collect dispatch channel with error: ", err.Error())
	}

	// Prepare queue

	q, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_FeedCollectDispatchQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Fatal("Failed to prepare MQ Feeds Collect dispatch queue with error: ", err.Error())
	}

	global.Logger.Debug("Feed Collect work start dispatching...")
	go func() {
		t := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-t.C:
				dispatchAllFeedCollectWorks(ch, q.Name)
			}
		}
	}()
}

func dispatchAllFeedCollectWorks(ch *amqp.Channel, queueName string) {
	global.Logger.Debug("Start dispatching feeds collect works...")

	nowTime := time.Now()

	// Accounts need update
	var accountsNeedUpdate []models.Account

	global.DB.Find(&accountsNeedUpdate, "next_update < ?", nowTime)

	if len(accountsNeedUpdate) == 0 {
		global.Logger.Debug("No accounts need update currently")
		return
	}

	// Dispatch update works
	for _, account := range accountsNeedUpdate {

		// Update account settings
		platform := commonConsts.SUPPORTED_PLATFORM[account.Platform]
		interv := nowTime.Sub(account.LastUpdated)
		if interv < platform.MinRefreshGap {
			interv = platform.MinRefreshGap
		} else if interv > platform.MaxRefreshGap {
			interv = platform.MaxRefreshGap
		}
		account.UpdateInterval = interv
		account.NextUpdate = nowTime.Add(interv)

		// Dispatch work
		work := commonTypes.WorkDispatched{
			DispatchAt: nowTime,
			AccountID:  account.ID,
			Platform:   account.Platform,
			Username:   account.Username,
			DropBefore: account.LastUpdated,
			DropAfter:  account.NextUpdate, // If cannot be performed before DDL, work fails (cause new work would replace current one)
		}

		if err := DispatchSingleFeedCollectWork(ch, &work, queueName); err != nil {
			global.Logger.Error("Failed to dispatch work: ", work)
		} else {
			// Update account
			global.DB.Save(&account)
		}

	}

	global.Logger.Debugf("Feeds collect works dispatched for %d accounts.", len(accountsNeedUpdate))
}

func DispatchSingleFeedCollectWork(ch *amqp.Channel, work *commonTypes.WorkDispatched, queueName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_FeedCollectDispatchTimeOut)
	defer cancel()

	if workBytes, err := json.Marshal(&work); err != nil {
		global.Logger.Error("Failed to marshall work: ", work)
		return err
	} else if err = ch.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/json",
			Body:         workBytes,
		},
	); err != nil {
		global.Logger.Error("Failed to dispatch work: ", work)
		return err
	} else {
		// Dispatched successfully
		global.Metrics.Work.Dispatched.Inc(1)
	}
	return nil
}
