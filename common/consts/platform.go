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
	"twitter": {
		Name:               "Twitter",
		FeedLink:           "{{rsshub_stateless}}/twitter/user/{{username}}/excludeReplies=1&includeRts=0",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
	"tg_channel": {
		Name:               "Telegram Channel",
		FeedLink:           "{{rsshub_stateless}}/telegram/channel/{{username}}/includeServiceMsg=0",
		MinRefreshGap:      30 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: false,
		HTML2Markdown:      true,
	},
	"substack": {
		Name:               "Substack",
		FeedLink:           "https://{{username}}.substack.com/feed",
		MinRefreshGap:      20 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: false,
		HTML2Markdown:      true,
	},
	"pixiv": {
		Name:               "pixiv",
		FeedLink:           "{{rsshub_stateful}}/pixiv/user/{{username}}", // Actually is UserID
		MinRefreshGap:      1 * time.Hour,
		MaxRefreshGap:      12 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
	"y2b_channel": {
		Name:               "YouTube Channel",
		FeedLink:           "https://www.youtube.com/feeds/videos.xml?channel_id={{username}}", // Channel ID
		MinRefreshGap:      1 * time.Hour,
		MaxRefreshGap:      12 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
	"mastodon": {
		Name:               "Mastodon",
		FeedLink:           "https://{{instance}}/@{{username}}.rss",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
	"jike": {
		Name:               "Jike",
		FeedLink:           "{{rsshub_stateless}}/jike/user/@{{username}}",
		MinRefreshGap:      10 * time.Minute,
		MaxRefreshGap:      1 * time.Hour,
		IsMediaAttachments: true,
		HTML2Markdown:      false,
	},
}
