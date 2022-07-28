package validateAccounts

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"strings"
)

func ValidateAccounts(m *nats.Msg) {
	global.Logger.Debug("New validate request received: ", string(m.Data))

	var validateReq commonTypes.ValidateRequest

	if err := json.Unmarshal(m.Data, &validateReq); err != nil {
		global.Logger.Error("Failed to parse validate request.", err.Error())
		handleFailed(m.Reply, commonConsts.ERROR_CODE_FAILED_TO_PARSE_JSON, "Failed to parse validate request")
		return
	}

	validateString := strings.ToLower(strings.ReplaceAll(commonConsts.ValidateTemplate, "{{crossbell_character_handle}}", validateReq.CrossbellCharacter))

	switch validateReq.Platform {
	case "medium":
		accountMedium(m.Reply, validateReq.Username, validateString)
	case "tiktok":
		accountTikTok(m.Reply, validateReq.Username, validateString)
	default:
		handleFailed(m.Reply, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
	}

}
