package medium

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"math"
	"regexp"
	"strings"
	"time"
)

var (
	imageRegex *regexp.Regexp
)

func init() {

	// Medium regex
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, acceptTime time.Time, collectLink string) {
	// Refer to https://medium.com/feed/@nya_9949

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for medium")

	rawFeed, errCode, err := utils.FeedRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		true,
	)
	if err != nil {
		callback.FeedsHandleFailed(work, acceptTime, errCode, err.Error())
		return
	}

	var feeds []commonTypes.RawFeed
	var minimalInterval time.Duration = math.MaxInt64

	for index, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Title:       item.Title,
				Link:        item.Link,
				GUID:        item.GUID,
				Categories:  item.Categories,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			rawContent := item.Content

			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			feed.Media = utils.UploadAllMedia(imgs)
			for _, media := range feed.Media {
				rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSURI)
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)
			if index > 0 {
				interv := rawFeed.Items[index-1].PublishedParsed.Sub(*item.PublishedParsed)
				if interv < 0 {
					interv = -interv
				}
				if interv < minimalInterval {
					minimalInterval = interv
				}
			}
		}
	}

	callback.FeedsHandleSucceeded(work, acceptTime, feeds, minimalInterval)

}
