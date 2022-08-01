package jobs

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"time"
)

func FeedCollectStartReceiveFailedWork() error {
	_, err := global.MQ.Subscribe(commonConsts.MQSETTINGS_FailedChannelName, feedCollectHandleFailed)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds Collect failed queue with error: ", err.Error())
		return err
	}

	return nil
}

func feedCollectHandleFailed(m *nats.Msg) {
	global.Logger.Warn("New failed Collect work received: ", string(m.Data))

	var workFailed commonTypes.WorkFailed
	if err := json.Unmarshal(m.Data, &workFailed); err != nil {
		global.Logger.Error("Unable to parse failed work: ", string(m.Data))
	} else {
		global.Logger.Warn("Work failed for: ", workFailed)
		global.Metrics.Work.Failed.Inc(1)

		// Retry work if still valid
		if time.Now().Before(workFailed.DropAfter) {
			_ = DispatchSingleFeedCollectWork(&workFailed.WorkDispatched)
		}
	}
}
