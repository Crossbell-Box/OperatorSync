package consts

import "time"

// `cos` means Crossbell Operator Sync
const (
	MQSETTINGS_PublishTimeOut = 5 * time.Second

	MQSETTINGS_FeedCollectDispatchQueueName  = "cos_feed_collect"
	MQSETTINGS_FeedCollectSucceededQueueName = "cos_feed_succeeded"
	MQSETTINGS_FeedCollectFailedQueueName    = "cos_feed_failed"

	MQSETTINGS_ValidateRPCQueueName   = "cos_validate"
	MQSETTINGS_ValidateRequestTimeOut = 10 * time.Second

	MQSETTINGS_OnChainRPCQueueName   = "cos_onchain"
	MQSETTINGS_OnChainRequestTimeOut = 3 * time.Minute
)
