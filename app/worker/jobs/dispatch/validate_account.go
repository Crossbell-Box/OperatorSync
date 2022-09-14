package dispatch

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
)

func ValidateAccounts(ch *amqp.Channel, d *amqp.Delivery) {
	global.Logger.Debug("New validate request received: ", string(d.Body))

	var validateReq commonTypes.ValidateRequest

	if err := json.Unmarshal(d.Body, &validateReq); err != nil {
		global.Logger.Error("Failed to parse validate request.", err.Error())
		callback.ValidateHandleFailed(ch, d, commonConsts.ERROR_CODE_FAILED_TO_PARSE_JSON, "Failed to parse validate request")
		return
	}

	handle, err := utils.GetCrossbellHandleFromID(validateReq.CrossbellCharacterID)
	if err != nil {
		callback.ValidateHandleFailed(ch, d, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error())
	}

	validateString := strings.ToLower(fmt.Sprintf("Crossbell@%s", handle))

	var (
		isSucceeded bool
		code        uint
		msg         string
		isValid     bool
	)

	switch validateReq.Platform {
	case "medium":
		isSucceeded, code, msg, isValid = medium.Account(validateReq.Username, validateString)
	case "tiktok":
		isSucceeded, code, msg, isValid = tiktok.Account(validateReq.Username, validateString)
	default:
		callback.ValidateHandleFailed(ch, d, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
		return
	}

	callback.ValidateHandleResponse(ch, d, isSucceeded, code, msg, isValid)

}
