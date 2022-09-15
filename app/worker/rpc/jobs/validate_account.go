package jobs

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strings"
)

func ValidateAccounts(validateReq *commonTypes.ValidateRequest) *commonTypes.ValidateResponse {
	global.Logger.Debug("New validate request received: ", validateReq)

	handle, err := utils.GetCrossbellHandleFromID(validateReq.CrossbellCharacterID)
	if err != nil {
		return ValidateHandleFailed(commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error())
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
		return ValidateHandleFailed(commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
	}

	return ValidateHandleResponse(isSucceeded, code, msg, isValid)

}

func ValidateHandleFailed(errCode uint, errMsg string) *commonTypes.ValidateResponse {
	return ValidateHandleResponse(
		false,
		errCode,
		errMsg,
		false,
	)

}

func ValidateHandleResponse(isSucceeded bool, code uint, msg string, isValid bool) *commonTypes.ValidateResponse {
	return &commonTypes.ValidateResponse{
		IsSucceeded:            isSucceeded,
		Code:                   code,
		Message:                msg,
		IsValidateStringExists: isValid,
	}
}
