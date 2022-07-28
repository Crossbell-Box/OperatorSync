package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/processFeeds"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/validateAccounts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
)

func StartProcessFeeds() error {

	cccs := &types.ConcurrencyChannels{
		Stateful:  types.NewCtrl(config.Config.ConcurrencyStateful),
		Stateless: types.NewCtrl(config.Config.ConcurrencyStateless),
		Direct:    types.NewCtrl(config.Config.ConcurrencyDirect),
	}

	for platformID, _ := range commonConsts.SUPPORTED_PLATFORM {
		channelName := commonConsts.MQSETTINGS_PlatformChannelPrefix + platformID
		if _, err := global.MQ.Subscribe(channelName, processFeeds.ProcessFeeds(platformID, cccs)); err != nil {
			global.Logger.Error("Failed to subscribe to MQ platform queue ", platformID, " with error: ", err.Error())
			return err
		}
	}

	return nil
}

func StartValidateAccounts() error {
	if _, err := global.MQ.Subscribe(commonConsts.MQSETTINGS_ValidateChannelName, validateAccounts.ValidateAccounts); err != nil {
		global.Logger.Error("Failed to subscribe to MQ validate queue with error: ", err.Error())
		return err
	}

	return nil
}
