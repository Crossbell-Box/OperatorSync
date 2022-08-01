package callback

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func FeedsHandleSucceeded(workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, rawFeeds []commonTypes.RawFeed, newInterval time.Duration) {

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
		err = global.MQ.Publish(commonConsts.MQSETTINGS_SucceededChannelName, succeededWorkBytes)
		if err != nil {
			global.Logger.Error("Failed to report succeeded work with error: ", err.Error())
		} else {
			global.Logger.Debug("And result pushed to succeeded channel.")
		}
	}
}

func FeedsHandleFailed(workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, errCode uint, errMsg string) {
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
		err = global.MQ.Publish(commonConsts.MQSETTINGS_FailedChannelName, failedWorkBytes)
		if err != nil {
			global.Logger.Error("Failed to report failed work with error: ", err.Error())
		}
	}
}
