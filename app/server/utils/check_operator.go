package utils

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func CheckOperator(crossbellCharacterID string) (bool, error) {

	// Why return true as default result?
	// Cause if check error, system might mistakenly
	// terminate users' accounts for false.

	checkRequest := commonTypes.CheckOperatorRequest{
		CrossbellCharacterID: crossbellCharacterID,
	}

	var checkResponse commonTypes.CheckOperatorResponse

	// Start request
	errChan := make(chan error, 1)

	go func() {
		errChan <- global.RPC.Call(
			fmt.Sprintf("%s.%s", commonConsts.RPCSETTINGS_BaseServiceName, commonConsts.RPCSETTINGS_CheckOperatorServiceName),
			checkRequest,
			&checkResponse,
		)
	}()

	// Set timeout
	select {
	case <-time.After(commonConsts.RPCSETTINGS_CheckOperatorRequestTimeOut):
		// Timeout
		global.Logger.Errorf("Check Operator request timeout...")
		return true, fmt.Errorf("check operator request timeout")

	case err := <-errChan:
		if err != nil {
			global.Logger.Errorf("Failed to receive operator check response with error: %s", err.Error())
			return true, err
		}
	}

	// Check operator response
	if !checkResponse.IsSucceeded {
		// Something is wrong
		global.Logger.Errorf("Check Operator work failed with error: %s", checkResponse.Message)
		return true, fmt.Errorf(checkResponse.Message)
	}

	return checkResponse.IsOperatorValid, nil

}
