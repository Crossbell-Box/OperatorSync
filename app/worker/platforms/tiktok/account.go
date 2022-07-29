package tiktok

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"strings"
)

func Account(mqReply string, username string, validateString string) {

	// RSSHub will index userinfo, so we need to parse their feeds
	// Sorry TikTok server :pray:

	if rawFeed, errCode, err := utils.RSSFeedRequest(
		strings.ReplaceAll(commonConsts.SUPPORTED_PLATFORM["tiktok"].FeedLink, "{{username}}", username),
		true,
	); err != nil {
		global.Logger.Error("Failed to collect tiktok feeds of user ", username, " for account validate with error: ", err.Error())
		callback.ValidateHandleFailed(mqReply, errCode, err.Error())
	} else {
		if strings.Contains(strings.ToLower(rawFeed.Description), validateString) {
			global.Logger.Debug("Account verify succeeded")
			callback.ValidateHandleSucceeded(mqReply, true)
		} else {
			global.Logger.Debug("No validate string found: ", rawFeed.Description)
			callback.ValidateHandleSucceeded(mqReply, false)
		}
	}

}
