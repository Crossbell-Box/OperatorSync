package consts

import "time"

// `cos` means Crossbell Operator Sync
const (
	MQSETTINGS_PlatformChannelPrefix = "cos_platform_"
	MQSETTINGS_SucceededChannelName  = "cos_succeeded"
	MQSETTINGS_FailedChannelName     = "cos_failed"

	MQSETTINGS_ValidateChannelName    = "validate"
	MQSETTINGS_ValidateRequestTimeOut = 3 * time.Second

	MQSETTINGS_OnChainStreamName          = "onchain"
	MQSETTINGS_OnChainDispatchSubjectName = "dispatch"
	MQSETTINGS_OnChainResponseSubjectName = "response"
)
