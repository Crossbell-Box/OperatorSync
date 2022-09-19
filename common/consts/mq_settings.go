package consts

import "time"

// `cos` means Crossbell Operator Sync
const (
	MQSETTINGS_PublishTimeOut = 5 * time.Second

	MQSETTINGS_FeedCollectDispatchQueueName   = "cos_q_feed_dispatch"
	MQSETTINGS_FeedCollectRetrieveQueueName   = "cos_q_feed_retrieve"
	MQSETTINGS_FeedCollectIdentifierField     = "cos-status-identifier"
	MQSETTINGS_FeedCollectSucceededIdentifier = "succeeded"
	MQSETTINGS_FeedCollectFailedIdentifier    = "failed"
)
