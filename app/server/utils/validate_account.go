package utils

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func ValidateAccount(crossbellCharacterID string, platform string, username string) (bool, error) {

	validateRequest := commonTypes.ValidateRequest{
		Platform:             platform,
		Username:             username,
		CrossbellCharacterID: crossbellCharacterID,
	}

	var validateResponse commonTypes.ValidateResponse

	// Start request
	errChan := make(chan error, 1)

	go func() {
		errChan <- global.RPC.Call(
			fmt.Sprintf("%s.%s", commonConsts.RPCSETTINGS_BaseServiceName, commonConsts.RPCSETTINGS_ValidateServiceName),
			validateRequest,
			&validateResponse,
		)
	}()

	// Set timeout
	select {
	case <-time.After(commonConsts.RPCSETTINGS_ValidateRequestTimeOut):
		// Timeout
		global.Logger.Errorf("Validate request timeout...")
		return false, fmt.Errorf("validate request timeout")

	case err := <-errChan:
		if err != nil {
			global.Logger.Errorf("Failed to receive validate response with error: %s", err.Error())
			return false, err
		}
	}

	// Validate response
	if !validateResponse.IsSucceeded {
		// Something is wrong
		global.Logger.Error("Validate work failed with code ", validateResponse.Code, " and message: ", validateResponse.Message)
		return false, fmt.Errorf(validateResponse.Message)
	}

	return validateResponse.IsValidateStringExists, nil
}
