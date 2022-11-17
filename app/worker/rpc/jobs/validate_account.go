package jobs

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/mastodon"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/pinterest"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/pixiv"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/substack"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tg_channel"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/twitter"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/y2b_channel"
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

	var accountValidateFunc func(string, string) (bool, uint, string, bool)

	switch validateReq.Platform {
	case "medium":
		accountValidateFunc = medium.Account
	case "tiktok":
		accountValidateFunc = tiktok.Account
	case "pinterest":
		accountValidateFunc = pinterest.Account
	case "twitter":
		accountValidateFunc = twitter.Account
	case "tg_channel":
		accountValidateFunc = tg_channel.Account
	case "substack":
		accountValidateFunc = substack.Account
	case "pixiv":
		accountValidateFunc = pixiv.Account
	case "y2b_channel":
		accountValidateFunc = y2b_channel.Account
	case "mastodon":
		accountValidateFunc = mastodon.Account
	default:
		ValidateHandleFailed(commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform", response)
		return
	}

	isSucceeded, code, msg, isValid = accountValidateFunc(validateReq.Username, validateString)

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
