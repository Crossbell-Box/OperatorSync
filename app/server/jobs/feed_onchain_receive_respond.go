package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
)

func FeedOnChainStartReceiveRespond() error {
	if _, err := global.MQJS.Subscribe(
		fmt.Sprintf("%s.%s", commonConsts.MQSETTINGS_OnChainStreamName, commonConsts.MQSETTINGS_OnChainResponseSubjectName),
		feedOnChainHandleRespond,
	); err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds OnChain respond queue with error: ", err.Error())
		return err
	}

	return nil
}

func feedOnChainHandleRespond(m *nats.Msg) {
	global.Logger.Debug("New OnChain respond received: ", string(m.Data))

	// Report received progress
	if err := m.InProgress(); err != nil {
		global.Logger.Error("Failed to report work progress")
	}

	// Parse response
	var workRespond commonTypes.OnChainRespond
	if err := json.Unmarshal(m.Data, &workRespond); err != nil {
		global.Logger.Error("Unable to parse respond: ", string(m.Data))
		if err := m.Nak(); err != nil {
			global.Logger.Error("Failed to report work Nak status")
		}
		return
	}

	// Validate response
	if workRespond.Platform == "" || workRespond.FeedID == 0 || workRespond.IPFSUri == "" {
		global.Logger.Error("Respond data is invalid")
		if err := m.Nak(); err != nil {
			global.Logger.Error("Failed to report work Nak status")
		}
		return
	}

	// Update feed from database
	platformSpecifiedFeed := models.Feed{
		Feed: types.Feed{
			Platform: workRespond.Platform,
		},
	}

	var feed models.Feed
	if err := global.DB.Scopes(models.FeedTable(platformSpecifiedFeed)).First(&feed, workRespond.FeedID).Error; err != nil {
		global.Logger.Error("Failed to find matching feed")
		if err := m.Nak(); err != nil {
			global.Logger.Error("Failed to report work Nak status")
		}
		return
	}
	feed.IPFSUri = workRespond.IPFSUri
	feed.Transaction = workRespond.Transaction
	if err := global.DB.Scopes(models.FeedTable(platformSpecifiedFeed)).Save(&feed).Error; err != nil {
		global.Logger.Error("Failed to save updated feed")
		if err := m.Nak(); err != nil {
			global.Logger.Error("Failed to report work Nak status")
		}
		return
	}

	if err := m.Ack(); err != nil {
		global.Logger.Error("Failed to report work Ack status")
	}
}
