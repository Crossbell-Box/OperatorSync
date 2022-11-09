package jobs

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/chain"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func CheckOperator(workDispatched *commonTypes.CheckOperatorRequest, response *commonTypes.CheckOperatorResponse) {
	result, err := chain.CheckOperator(workDispatched.CrossbellCharacterID)
	if err != nil {
		*response = commonTypes.CheckOperatorResponse{
			IsSucceeded:     false,
			Message:         err.Error(),
			IsOperatorValid: false,
		}
	} else {
		*response = commonTypes.CheckOperatorResponse{
			IsSucceeded:     true,
			Message:         "",
			IsOperatorValid: result,
		}
	}
}
