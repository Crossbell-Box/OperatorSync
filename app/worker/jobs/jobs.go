package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/dispatch"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
)

func FeedCollectStartProcess() error {

	cccs := &types.ConcurrencyChannels{
		Stateful:  types.NewCtrl(config.Config.ConcurrencyStateful),
		Stateless: types.NewCtrl(config.Config.ConcurrencyStateless),
		Direct:    types.NewCtrl(config.Config.ConcurrencyDirect),
	}

	for platformID, _ := range commonConsts.SUPPORTED_PLATFORM {
		channelName := commonConsts.MQSETTINGS_PlatformChannelPrefix + platformID
		if _, err := global.MQ.Subscribe(channelName, dispatch.ProcessFeeds(platformID, cccs)); err != nil {
			global.Logger.Error("Failed to subscribe to MQ platform queue ", platformID, " with error: ", err.Error())
			return err
		}
	}

	return nil
}

func AccountValidateStartProcess() error {
	if _, err := global.MQ.Subscribe(commonConsts.MQSETTINGS_ValidateChannelName, dispatch.ValidateAccounts); err != nil {
		global.Logger.Error("Failed to subscribe to MQ validate queue with error: ", err.Error())
		return err
	}

	return nil
}

func OnChainFeedsStartProcess() error {
	if _, err := global.MQ.Subscribe(commonConsts.MQSETTINGS_OnChainChannelName, dispatch.OnChain); err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds OnChain dispatch queue with error: ", err.Error())
		return err
	}

	return nil
}
