package callback

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func ValidateHandleSucceeded(mqReply string, isValid bool) {
	if succeededResponseBytes, err := json.Marshal(&commonTypes.ValidateResponse{
		IsSucceeded:            true,
		Code:                   0,
		Message:                "",
		IsValidateStringExists: isValid,
	}); err != nil {
		global.Logger.Errorf("Failed to marshall succeeded validate request with error: %s", err.Error())
	} else {
		if err := global.MQ.Publish(mqReply, succeededResponseBytes); err != nil {
			global.Logger.Error("Failed to report succeeded validate request with error: ", err.Error())
		}
	}
}

func ValidateHandleFailed(mqReply string, errCode uint, errMsg string) {
	failedResponse := commonTypes.ValidateResponse{
		IsSucceeded:            false,
		Code:                   errCode,
		Message:                errMsg,
		IsValidateStringExists: false,
	}

	if failedResponseBytes, err := json.Marshal(&failedResponse); err != nil {
		global.Logger.Error("Failed to marshall failed validate request: ", failedResponse)
	} else {
		err = global.MQ.Publish(mqReply, failedResponseBytes)
		if err != nil {
			global.Logger.Error("Failed to report failed validate request with error: ", err.Error())
		}
	}

}
