package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func ValidateAccount(crossbellCharacterID string, platform string, username string) (bool, error) {
	// Convert crossbell character id to handle

	var err error

	validateRequest := commonTypes.ValidateRequest{
		Platform:             platform,
		Username:             username,
		CrossbellCharacterID: crossbellCharacterID,
	}

	var validateRequestBytes []byte

	if validateRequestBytes, err = json.Marshal(&validateRequest); err != nil {
		global.Logger.Error("Failed to parse validate request to bytes: ", err.Error())
		return false, err
	}

	// Request validate

	validateResponseMsg, err := global.MQ.Request(commonConsts.MQSETTINGS_ValidateChannelName, validateRequestBytes, commonConsts.MQSETTINGS_ValidateRequestTimeOut)
	if err != nil {
		global.Logger.Error("Failed to get validate response", err.Error())
		return false, err
	}

	var validateResponse commonTypes.ValidateResponse
	if err := json.Unmarshal(validateResponseMsg.Data, &validateResponse); err != nil {
		global.Logger.Error("Failed to parse validate response", err.Error())
		return false, err
	}

	if !validateResponse.IsSucceeded {
		// Something is wrong
		global.Logger.Error("Validate work failed with code ", validateResponse.Code, " and message: ", validateResponse.Message)
		return false, fmt.Errorf(validateResponse.Message)
	}

	return validateResponse.IsValidateStringExists, nil
}
