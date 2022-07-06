package consts

import "time"

type platform = struct {
	Name           string
	FeedLink       string
	FeedIsStateful bool
	MinRefreshGap  time.Duration
	MaxRefreshGap  time.Duration
}

// FeedLink replace rule:
// - Replace {{username}} with real username
// - Replace {{rsshub_stateful}} with Stateful (Logged in) RSSHub address
// - Replace {{rsshub_stateless}} with Stateless (Not logged in) RSSHub address

var SUPPORTED_PLATFORM = map[string]platform{
	"medium": {
		Name:           "Medium",
		FeedLink:       "https://medium.com/feed/@{{username}}",
		FeedIsStateful: false,
		MinRefreshGap:  12 * time.Hour,     // Half of a day
		MaxRefreshGap:  7 * 24 * time.Hour, // One week
	},
	"tiktok": {
		Name:           "TikTok",
		FeedLink:       "{{rsshub_stateful}}/tiktok/user/@{{username}}",
		FeedIsStateful: true,
		MinRefreshGap:  1 * time.Hour,  // 1 Hour
		MaxRefreshGap:  24 * time.Hour, // 1 Day
	},
}