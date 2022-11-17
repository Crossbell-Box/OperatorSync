package mastodon

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	"go.uber.org/zap"
	"testing"
)

func TestRawFeed(t *testing.T) {
	// Init deps
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	collectLink := "https://thenomads.social/@candinya.rss"

	rawFeed, errCode, err := utils.RSSFeedRequest(
		collectLink,
		false,
	)

	medias, ok := rawFeed.Items[1].Extensions["media"]["content"]
	if ok {
		for _, m := range medias {
			t.Log(m.Attrs["url"])
		}
	} else {
		t.Log("No medias")
	}

	t.Log(rawFeed, errCode, err)
}
