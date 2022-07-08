package consts

import "time"

const (
	CACHE_PREFIX               = "cos:cache"
	MAX_FEED_LIMIT_PER_PAGE    = 200
	MAX_MEDIA_LIMIT_PER_PAGE   = 200
	ACCOUNT_LIST_CACHE_EXPIRE  = 10 * time.Minute
	ACCOUNT_REFRESH_COOLDOWN   = 1 * time.Minute
	WORK_COUNTS_METRICS_PREFIX = "work_counts_"
)
