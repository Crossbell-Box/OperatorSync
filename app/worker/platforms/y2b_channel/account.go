package y2b_channel

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"regexp"
	"strings"
)

var (
	y2bChannelDescriptionRegex *regexp.Regexp
)

func init() {
	y2bChannelDescriptionRegex = regexp.MustCompile(`<meta[^>]+\bname="description"[^>]+\bcontent=["']([^"']+)["']`)
}

func Account(username string, validateString string) (bool, uint, string, bool) {

	global.Logger.Debug("Validate string: ", validateString)

	// Get HTTP Page
	pageContent, err := utils.HttpRequest(fmt.Sprintf("https://www.youtube.com/channel/%s", username), true)
	if err != nil {
		global.Logger.Error("Failed to check youtube page of channel ", username, " for account validate with error: ", err.Error())
		return false, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error(), false
	} else {
		description := y2bChannelDescriptionRegex.FindStringSubmatch(string(pageContent))
		if len(description) < 2 {
			global.Logger.Error("Failed to check youtube page of channel ", username, " : failed to find description string")
			return false, commonConsts.ERROR_CODE_FAILED_TO_FIND_NECESSARY_FIELD, "Failed to find description string", false
		} else {
			if strings.Contains(strings.ToLower(description[1]), validateString) {
				global.Logger.Debug("Account verify succeeded")
				return true, 0, "", true
			} else {
				global.Logger.Debug("No validate string found: ", description[1])
				return true, 0, "", false
			}
		}

	}
}
