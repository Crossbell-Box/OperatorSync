package consts

import "time"

type platform = struct {
	// Meta information
	Name          string
	FeedLink      string
	MinRefreshGap time.Duration
	MaxRefreshGap time.Duration

	// OnChain Settings
	IsMediaAttachments bool
	HTML2Markdown      bool
}

// FeedLink replace rule:
// - Replace {{username}} with real username
// - Replace {{rsshub_stateful}} with Stateful (Logged in) RSSHub address
// - Replace {{rsshub_stateless}} with Stateless (Not logged in) RSSHub address

var SUPPORTED_PLATFORM = map[string]platform{
	"medium": {
		Name:               "Medium",
		FeedLink:           "https://medium.com/feed/@{{username}}",
		MinRefreshGap:      20 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: false,
		HTML2Markdown:      true,
	},
	"tiktok": {
		Name:               "TikTok",
		FeedLink:           "{{rsshub_stateless}}/tiktok/user/@{{username}}",
		MinRefreshGap:      30 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
	"pinterest": {
		Name:               "Pinterest",
		FeedLink:           "https://www.pinterest.com/{{username}}/feed.rss",
		MinRefreshGap:      30 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
}
