package y2b_channel

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	"testing"
)

func TestParseFeeds(t *testing.T) {
	rawFeed, _, _ := utils.RSSFeedRequest(
		"https://www.youtube.com/feeds/videos.xml?channel_id=UCs_f32UpeFk2CM8hzF7Q61Q",
		false,
	)
	t.Log(rawFeed)
	t.Log(rawFeed.Items[0].Extensions["media"]["group"][0].Children["description"][0].Value)
}
