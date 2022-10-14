package callback

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func FeedsHandleSucceeded(ch *amqp.Channel, qSucceededName string, workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, rawFeeds []commonTypes.RawFeed, newInterval time.Duration) {

	global.Logger.Debug("Work succeeded: ", workDispatched)

	if len(rawFeeds) == 0 {
		// Actually nothing
		global.Logger.Debug("... but nothing found")
	}

	succeededWork := commonTypes.WorkSucceeded{
		WorkDispatched: *workDispatched,
		AcceptedAt:     acceptTime,
		SucceededAt:    time.Now(),
		Feeds:          rawFeeds,
		NewInterval:    newInterval,
	}
	if succeededWorkBytes, err := json.Marshal(&succeededWork); err != nil {
		global.Logger.Error("Failed to marshall succeeded work: ", succeededWork)
	} else {
		// Save result to redis
		storageKey := fmt.Sprintf(commonConsts.REDIS_FeedCollectResultKeyTemplate, workDispatched.Platform, workDispatched.Username, workDispatched.DispatchAt.UnixNano())
		commonGlobal.Redis.Set(context.Background(), storageKey, succeededWorkBytes, commonConsts.REDIS_FeedCollectResultExpires)

		ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_PublishTimeOut)
		defer cancel()

		err = ch.PublishWithContext(
			ctx,
			"",
			qSucceededName,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Headers: amqp.Table{
					commonConsts.MQSETTINGS_FeedCollectIdentifierField: commonConsts.MQSETTINGS_FeedCollectSucceededIdentifier,
				},
				Body: []byte(storageKey),
			},
		)
		if err != nil {
			global.Logger.Error("Failed to report succeeded work with error: ", err.Error())
		} else {
			global.Logger.Debug("And result pushed to succeeded channel.")
		}
	}
}

func FeedsHandleFailed(ch *amqp.Channel, qFailedName string, workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, errCode uint, errMsg string) {
	global.Logger.Warn("Work failed: ", workDispatched, " with code: ", errCode, " , reason: ", errMsg)

	failedWork := commonTypes.WorkFailed{
		WorkDispatched: *workDispatched,
		AcceptAt:       acceptTime,
		ErrorAt:        time.Now(),
		ErrorCode:      errCode,
		ErrorReason:    errMsg,
	}
	if failedWorkBytes, err := json.Marshal(&failedWork); err != nil {
		global.Logger.Error("Failed to marshall failed work: ", failedWork)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_PublishTimeOut)
		defer cancel()

		err = ch.PublishWithContext(
			ctx,
			"",
			qFailedName,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Headers: amqp.Table{
					commonConsts.MQSETTINGS_FeedCollectIdentifierField: commonConsts.MQSETTINGS_FeedCollectFailedIdentifier,
				},
				Body: failedWorkBytes,
			},
		)
		if err != nil {
			global.Logger.Error("Failed to report failed work with error: ", err.Error())
		}
	}
}
