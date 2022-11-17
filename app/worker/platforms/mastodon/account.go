package mastodon

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonUtils "github.com/Crossbell-Box/OperatorSync/common/utils"
	"net/http"
	"strings"
)

type ActivityPubAccountResponse struct {
	Summary string `json:"summary"`
	// Ignore others

	// Or error
	Error string `json:"error"`
}

func Account(usernameAtInstance string, validateString string) (bool, uint, string, bool) {

	global.Logger.Debug("Validate string: ", validateString)

	username, instance, err := commonUtils.SplitFediverseUsernameInstance(usernameAtInstance)
	if err != nil {
		global.Logger.Errorf("Failed to extract username & instance from proviede string %s with error: %s", usernameAtInstance, err.Error())
		return false, commonConsts.ERROR_CODE_INVALID_FORMAT, err.Error(), false
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/@%s", instance, username), nil)
	if err != nil {
		global.Logger.Errorf("Failed to prepare request with error: %s", err.Error())
		return false, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error(), false
	}

	req.Header.Set("Accept", "application/ld+json, application/json") // To get json response

	var res ActivityPubAccountResponse

	resEntity, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to finish request with error: %s", err.Error())
		return false, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error(), false
	}

	err = json.NewDecoder(resEntity.Body).Decode(&res)
	if err != nil {
		global.Logger.Errorf("Failed to parse response with error: %s", err.Error())
		return false, commonConsts.ERROR_CODE_FAILED_TO_PARSE_JSON, err.Error(), false
	}

	if res.Error != "" && res.Summary == "" {
		// Failed
		global.Logger.Errorf("Failed to find validate string with error: %s", err.Error())
		return true, commonConsts.ERROR_CODE_FAILED_TO_FIND_NECESSARY_FIELD, err.Error(), false
	} else if strings.Contains(strings.ToLower(res.Summary), validateString) {
		global.Logger.Debug("Account verify succeeded")
		return true, 0, "", true
	} else {
		global.Logger.Debug("No validate string found: ", res.Summary)
		return true, 0, "", false
	}

}
