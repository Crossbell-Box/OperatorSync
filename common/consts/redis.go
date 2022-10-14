package consts

import "time"

const (
	REDIS_FeedCollectResultKeyTemplate = "cos:com:%s:%s:%d" // platform : username : timestamp (work dispatched)
	REDIS_FeedCollectResultExpires     = 1 * time.Hour
)
