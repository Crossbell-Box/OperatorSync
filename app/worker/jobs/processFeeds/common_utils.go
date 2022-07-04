package processFeeds

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/mmcdole/gofeed"
	"time"
)

func makeRequest(url string, withProxy bool) ([]*gofeed.Item, uint, error) {
	feedsBody, err := utils.HttpRequest(url, withProxy)
	if err != nil {
		return nil, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(feedsBody))
	if err != nil {
		return nil, commonConsts.ERROR_CODE_FAILED_TO_PARSE_FEEDS, err
	}

	return feed.Items, 0, nil
}

func handleSucceeded(workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, rawFeeds []commonTypes.RawFeed, newInterval time.Duration) {
	if len(rawFeeds) == 0 {
		// Actually nothing
		return
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
		}
	}
}

func handleFailed(workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, errCode uint, errMsg string) {
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
