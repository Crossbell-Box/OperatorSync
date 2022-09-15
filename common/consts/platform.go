package consts

import "time"

type platform = struct {
	Name               string
	FeedLink           string
	MinRefreshGap      time.Duration
	MaxRefreshGap      time.Duration
	IsMediaAttachments bool
}

// FeedLink replace rule:
// - Replace {{username}} with real username
// - Replace {{rsshub_stateful}} with Stateful (Logged in) RSSHub address
// - Replace {{rsshub_stateless}} with Stateless (Not logged in) RSSHub address

var SUPPORTED_PLATFORM = map[string]platform{
	"medium": {
		Name:               "Medium",
		FeedLink:           "https://medium.com/feed/@{{username}}",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      3 * time.Hour,
		IsMediaAttachments: false,
	},
	"tiktok": {
		Name:               "TikTok",
		FeedLink:           "{{rsshub_stateless}}/tiktok/user/@{{username}}",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      3 * time.Hour,
		IsMediaAttachments: true,
	},
}
