package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/nats-io/nats.go"
)

func Jobs() error {

	// Start listening to FeedCollect works response
	if err := jobs.FeedCollectStartReceiveSucceededWork(); err != nil {
		return err
	}
	if err := jobs.FeedCollectStartReceiveFailedWork(); err != nil {
		return err
	}

	if config.Config.IsMainServer {
		// Check JetStream
		createOnChainStreamIfNonExist()
	}

	// Start listening to FeedOnChain works response
	if err := jobs.FeedOnChainStartReceiveRespond(); err != nil {
		return err
	}

	if config.Config.IsMainServer {
		// Start dispatch flush works
		jobs.FeedCollectStartDispatchWork()
	}

	return nil

}

func createOnChainStreamIfNonExist() {
	// Try to create main on-chain stream
	info, err := global.MQJS.StreamInfo(commonConsts.MQSETTINGS_OnChainStreamName)
	if err != nil {
		global.Logger.Errorf("Failed to get stream info: %v", err)
	}
	if info == nil {
		// Non-exist
		global.Logger.Debug("Creating feeds OnChain stream channel...")
		if _, err := global.MQJS.AddStream(&nats.StreamConfig{
			Name:        commonConsts.MQSETTINGS_OnChainStreamName,
			Description: "Channel for feeds collected that needs to be published on chain",
			Subjects: []string{
				fmt.Sprintf("%s.*", commonConsts.MQSETTINGS_OnChainStreamName),
			},
			Retention: nats.WorkQueuePolicy,
			Discard:   nats.DiscardNew,
			Storage:   nats.FileStorage,
		}); err != nil {
			global.Logger.Errorf("Failed to create stream: %v", err)
		}
	}
}
