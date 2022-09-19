package medium

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"regexp"
	"strings"
)

var (
	mediumDescriptionRegex *regexp.Regexp
)

func init() {
	mediumDescriptionRegex = regexp.MustCompile(`<meta[^>]+\bname="description"[^>]+\bcontent=["']([^"']+)["']`)
}

func Account(username string, validateString string) (bool, uint, string, bool) {

	global.Logger.Debug("Validate string: ", validateString)

	// Get HTTP Page
	pageContent, err := utils.HttpRequest(fmt.Sprintf("https://medium.com/@%s", username), true)
	if err != nil {
		global.Logger.Error("Failed to check medium page of user ", username, " for account validate with error: ", err.Error())
		return false, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error(), false
	} else {
		description := mediumDescriptionRegex.FindStringSubmatch(string(pageContent))
		if len(description) < 2 {
			global.Logger.Error("Failed to check medium page of user ", username, " : failed to find description string")
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
