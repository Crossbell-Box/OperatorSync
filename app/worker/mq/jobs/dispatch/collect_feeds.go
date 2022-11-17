package dispatch

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/mq/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/mastodon"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/pinterest"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/pixiv"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/substack"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tg_channel"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/twitter"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/y2b_channel"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func ProcessFeeds(cccs *types.ConcurrencyChannels, ch *amqp.Channel, qRetrieveName string, d *amqp.Delivery) {

	global.Logger.Debug("New work received: ", string(d.Body))

	acceptTime := time.Now()

	var workDispatched commonTypes.WorkDispatched

	if err := json.Unmarshal(d.Body, &workDispatched); err != nil {
		global.Logger.Error("Failed to parse dispatched work data.", err.Error())
		// Even unable to report to failed works cause work cannot be parsed
		return
	}

	collectLink := commonConsts.SUPPORTED_PLATFORM[workDispatched.Platform].FeedLink

	var (
		isSucceeded bool
		feeds       []commonTypes.RawFeed
		errCode     uint
		errMsg      string
	)

	var feedCollectFunc func(*types.ConcurrencyChannels, *commonTypes.WorkDispatched, string) (bool, []commonTypes.RawFeed, uint, string)

	switch workDispatched.Platform {
	case "medium":
		feedCollectFunc = medium.Feeds
	case "tiktok":
		feedCollectFunc = tiktok.Feeds
	case "pinterest":
		feedCollectFunc = pinterest.Feeds
	case "twitter":
		feedCollectFunc = twitter.Feeds
	case "tg_channel":
		feedCollectFunc = tg_channel.Feeds
	case "substack":
		feedCollectFunc = substack.Feeds
	case "pixiv":
		feedCollectFunc = pixiv.Feeds
	case "y2b_channel":
		feedCollectFunc = y2b_channel.Feeds
	case "mastodon":
		feedCollectFunc = mastodon.Feeds
	default:
		// Unable to handle
		callback.FeedsHandleFailed(ch, qRetrieveName, &workDispatched, acceptTime, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
		return
	}

	isSucceeded, feeds, errCode, errMsg = feedCollectFunc(cccs, &workDispatched, collectLink)

	if isSucceeded {
		callback.FeedsHandleSucceeded(ch, qRetrieveName, &workDispatched, acceptTime, feeds, utils.CalcNewInterval(&workDispatched, feeds))
	} else {
		callback.FeedsHandleFailed(ch, qRetrieveName, &workDispatched, acceptTime, errCode, errMsg)
	}
}
