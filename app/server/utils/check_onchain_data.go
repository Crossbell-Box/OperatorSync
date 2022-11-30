package utils

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func CheckOnChainData(crossbellCharacterID string) (bool, []commonTypes.Account, error) {

	// Why return true as default result?
	// Cause if check error, system might mistakenly
	// terminate users' accounts for false.

	checkRequest := commonTypes.CheckOnChainDataRequest{
		CrossbellCharacterID: crossbellCharacterID,
	}

	var checkResponse commonTypes.CheckOnChainDataResponse

	// Start request
	errChan := make(chan error, 1)

	go func() {
		errChan <- global.RPC.Call(
			fmt.Sprintf("%s.%s", commonConsts.RPCSETTINGS_BaseServiceName, commonConsts.RPCSETTINGS_CheckOnChainDataServiceName),
			checkRequest,
			&checkResponse,
		)
	}()

	// Set timeout
	select {
	case <-time.After(commonConsts.RPCSETTINGS_CheckOnChainDataRequestTimeOut):
		// Timeout
		global.Logger.Errorf("Check Operator request timeout...")
		return true, nil, fmt.Errorf("check operator request timeout")

	case err := <-errChan:
		if err != nil {
			global.Logger.Errorf("Failed to receive operator check response with error: %s", err.Error())
			return true, nil, err
		}
	}

	// Check on-chain data response
	if !checkResponse.IsSucceeded {
		// Something is wrong
		global.Logger.Errorf("Check OnChain data work failed with error: %s", checkResponse.Message)
		return true, nil, fmt.Errorf(checkResponse.Message)
	}

	return checkResponse.IsOperatorValid, checkResponse.ConnectedAccounts, nil

}
