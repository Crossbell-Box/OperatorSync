package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/processFeeds"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
)

func StartProcessFeeds() error {
	for platformID, _ := range commonConsts.SUPPORTED_PLATFORM {
		channelName := commonConsts.MQSETTINGS_PlatformChannelPrefix + platformID
		if _, err := global.MQ.Subscribe(channelName, processFeeds.ProcessFeeds(platformID)); err != nil {
			global.Logger.Error("Failed to subscribe to MQ platform queue ", platformID, " with error: ", err.Error())
			return err
		}
	}

	//defer sub.Drain() // Ignore errors

	return nil
}
