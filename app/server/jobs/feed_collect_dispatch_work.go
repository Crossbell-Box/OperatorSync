package jobs

import (
	"context"
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
	"time"
)

func FeedCollectStartDispatchWork() {

	// Start queue

	global.Logger.Debug("Feed Collect work start dispatching...")
	go func() {
		t := time.NewTicker(consts.JOBS_INTERVAL_FEED_COLLECT)
		for {
			// Waiting for connection
			for {
				time.Sleep(commonConsts.MQSETTINGS_ReconnectDelay)
				if commonGlobal.MQ != nil {
					break
				}
			}
			notifyClose := make(chan *amqp.Error)
			ch, q := prepareFeedCollectQueue(commonConsts.MQSETTINGS_FeedCollectDispatchQueueName, notifyClose)
			isPendingRestart := false
			for {
				if isPendingRestart {
					break
				}
				select {
				case <-t.C:
					go dispatchAllFeedCollectWorks(ch, q.Name)
				case err := <-notifyClose:
					if err != nil {
						global.Logger.Errorf("MQ channel closed with error %d (%s), preparing to reconnect", err.Code, err.Error())
						isPendingRestart = true
					}
				}
			}
		}
	}()
}

var (
	_isFeedCollectDispatchWorkProcessing bool
)

func init() {
	_isFeedCollectDispatchWorkProcessing = false
}

func dispatchAllFeedCollectWorks(ch *amqp.Channel, queueName string) {

	if _isFeedCollectDispatchWorkProcessing {
		global.Logger.Warn("Another FeedCollectDispatch work is running, skip this.")
		return
	}

	// Set busy flag
	_isFeedCollectDispatchWorkProcessing = true
	defer func() {
		_isFeedCollectDispatchWorkProcessing = false
	}()

	global.Logger.Debug("Start dispatching feeds collect works...")

	config.Status.Jobs.FeedCollectLastRun = time.Now()

	if config.Config.HeartBeatWebhooks.FeedCollect != "" {
		// Send heartbeat packet
		global.Logger.Debug("Sending feed collect heartbeat packet")
		_, err := (&http.Client{}).Get(config.Config.HeartBeatWebhooks.FeedCollect) // Ignore response
		if err != nil {
			global.Logger.Errorf("Failed to send feed collect heartbeat packet with error: %s", err.Error())
		}
	}

	nowTime := time.Now()

	// Accounts need update
	var accountsNeedUpdate []models.Account

	global.DB.Find(&accountsNeedUpdate, "next_update < ?", nowTime)

	if len(accountsNeedUpdate) == 0 {
		global.Logger.Debug("No accounts need update currently")
		return
	}

	global.Logger.Debugf("Found %d accounts need update", len(accountsNeedUpdate))

	// Dispatch update works
	for _, account := range accountsNeedUpdate {

		global.Logger.Debugf("Dispatching account update work for #%d (%s@%s)", account.ID, account.Username, account.Platform)

		// Update account settings
		platform, ok := commonConsts.SUPPORTED_PLATFORM[account.Platform]
		if !ok {
			// Unsupported platform
			// Might be caused by old feed collect work returns
			// Disable account
			global.DB.Delete(&account)
			continue
		}
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

			global.Logger.Debugf("Account update work for #%d (%s@%s) dispatched successfully", account.ID, account.Username, account.Platform)
		}

	}

	global.Logger.Debugf("Feeds collect works dispatched for %d accounts.", len(accountsNeedUpdate))
}

func DispatchSingleFeedCollectWork(ch *amqp.Channel, work *commonTypes.WorkDispatched, queueName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_PublishTimeOut)
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
			ContentType:  "application/json",
			Body:         workBytes,
		},
	); err != nil {
		global.Logger.Errorf("Failed to dispatch work %v with error %s", work, err.Error())
		return err
	}
	return nil
}
