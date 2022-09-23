package consts

import "time"

// `cos` means Crossbell Operator Sync
const (
	RPCSETTINGS_BaseServiceName = "COSWorkerV2"

	RPCSETTINGS_ValidateServiceName    = "Validate" // Should be same as function name
	RPCSETTINGS_ValidateRequestTimeOut = 20 * time.Second

	RPCSETTINGS_OnChainServiceName    = "OnChain" // Should be same as function name
	RPCSETTINGS_OnChainRequestTimeOut = 3 * time.Minute
)
