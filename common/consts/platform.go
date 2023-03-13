package consts

import "time"

type platform = struct {
	// Meta information
	Name          string
	FeedLink      string
	MinRefreshGap time.Duration
	MaxRefreshGap time.Duration
	Limit1Account bool

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
		Limit1Account:      true,
	},
	"tiktok": {
		Name:               "TikTok",
		FeedLink:           "{{rsshub_stateless}}/tiktok/user/@{{username}}",
		MinRefreshGap:      30 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
	"pinterest": {
		Name:               "Pinterest",
		FeedLink:           "https://www.pinterest.com/{{username}}/feed.rss",
		MinRefreshGap:      30 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
	"twitter": {
		Name:               "Twitter",
		FeedLink:           "{{rsshub_stateless}}/twitter/user/{{username}}/excludeReplies=0&includeRts=1&showSymbolForRetweetAndReply=false.json",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
	"tg_channel": {
		Name:               "Telegram Channel",
		FeedLink:           "{{rsshub_stateless}}/telegram/channel/{{username}}/includeServiceMsg=0.json",
		MinRefreshGap:      30 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: false,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
	"substack": {
		Name:               "Substack",
		FeedLink:           "https://{{username}}.substack.com/feed",
		MinRefreshGap:      20 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: false,
		HTML2Markdown:      true,
		Limit1Account:      true,
	},
	"pixiv": {
		Name:               "pixiv",
		FeedLink:           "{{rsshub_stateful}}/pixiv/user/{{username}}", // Actually is UserID
		MinRefreshGap:      1 * time.Hour,
		MaxRefreshGap:      12 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
	"y2b_channel": {
		Name:               "YouTube Channel",
		FeedLink:           "https://www.youtube.com/feeds/videos.xml?channel_id={{username}}", // Channel ID
		MinRefreshGap:      1 * time.Hour,
		MaxRefreshGap:      12 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
	"mastodon": {
		Name:               "Mastodon",
		FeedLink:           "https://{{instance}}/@{{username}}.rss",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      false,
	},
	"jike": {
		Name:               "Jike",
		FeedLink:           "{{rsshub_stateless}}/jike/user/{{username}}.json",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
		Limit1Account:      true,
	},
}
