package jobs

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/pinterest"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strings"
)

func ValidateAccounts(validateReq *commonTypes.ValidateRequest, response *commonTypes.ValidateResponse) {
	global.Logger.Debug("New validate request received: ", validateReq)

	handle, err := utils.GetCrossbellHandleFromID(validateReq.CrossbellCharacterID)
	if err != nil {
		ValidateHandleFailed(commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error(), response)
	}

	validateString := strings.ToLower(fmt.Sprintf("%s@Crossbell", handle))

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
	case "pinterest":
		isSucceeded, code, msg, isValid = pinterest.Account(validateReq.Username, validateString)
	default:
		ValidateHandleFailed(commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform", response)
	}

	ValidateHandleResponse(isSucceeded, code, msg, isValid, response)

}

func ValidateHandleFailed(errCode uint, errMsg string, response *commonTypes.ValidateResponse) {
	ValidateHandleResponse(
		false,
		errCode,
		errMsg,
		false,
		response,
	)

}

func ValidateHandleResponse(isSucceeded bool, code uint, msg string, isValid bool, response *commonTypes.ValidateResponse) {
	*response = commonTypes.ValidateResponse{
		IsSucceeded:            isSucceeded,
		Code:                   code,
		Message:                msg,
		IsValidateStringExists: isValid,
	}
}
