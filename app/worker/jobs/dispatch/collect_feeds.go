package dispatch

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
	"time"
)

func ProcessFeeds(cccs *types.ConcurrencyChannels, ch *amqp.Channel, qSucceededName string, qFailedName string, d *amqp.Delivery) {

	global.Logger.Debug("New work received: ", string(d.Body))

	acceptTime := time.Now()

	var workDispatched commonTypes.WorkDispatched

	if err := json.Unmarshal(d.Body, &workDispatched); err != nil {
		global.Logger.Error("Failed to parse dispatched work data.", err.Error())
		// Even unable to report to failed works cause work cannot be parsed
		return
	}

	collectLink := commonConsts.SUPPORTED_PLATFORM[workDispatched.Platform].FeedLink

	collectLink =
		strings.ReplaceAll(
			strings.ReplaceAll(
				collectLink,
				"{{rsshub_stateful}}",
				config.Config.RSSHubEndpointStateful,
			),
			"{{rsshub_stateless}}",
			config.Config.RSSHubEndpointStateless,
		)

	var (
		isSucceeded bool
		feeds       []commonTypes.RawFeed
		newInterval time.Duration
		errCode     uint
		errMsg      string
	)

	switch workDispatched.Platform {
	case "medium":
		isSucceeded, feeds, newInterval, errCode, errMsg = medium.Feeds(cccs, &workDispatched, collectLink)
	case "tiktok":
		isSucceeded, feeds, newInterval, errCode, errMsg = tiktok.Feeds(cccs, &workDispatched, collectLink)
	default:
		// Unable to handle
		callback.FeedsHandleFailed(ch, qFailedName, &workDispatched, acceptTime, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
		return
	}

	if isSucceeded {
		callback.FeedsHandleSucceeded(ch, qSucceededName, &workDispatched, acceptTime, feeds, newInterval)
	} else {
		callback.FeedsHandleFailed(ch, qFailedName, &workDispatched, acceptTime, errCode, errMsg)
	}
}
