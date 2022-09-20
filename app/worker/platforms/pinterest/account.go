package pinterest

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"strings"
)

func Account(username string, validateString string) (bool, uint, string, bool) {

	global.Logger.Debug("Validate string: ", validateString)

	collectLink := commonConsts.SUPPORTED_PLATFORM["pinterest"].FeedLink

	if rawFeed, errCode, err := utils.RSSFeedRequest(
		strings.ReplaceAll(collectLink, "{{username}}", username),
		true,
	); err != nil {
		global.Logger.Error("Failed to collect pinterest feeds of user ", username, " for account validate with error: ", err.Error())
		return false, errCode, err.Error(), false
	} else {
		if strings.Contains(strings.ToLower(rawFeed.Description), validateString) {
			global.Logger.Debug("Account verify succeeded")
			return true, 0, "", true
		} else {
			global.Logger.Debug("No validate string found: ", rawFeed.Description)
			return true, 0, "", false
		}
	}

}
