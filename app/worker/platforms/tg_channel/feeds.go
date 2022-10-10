package tg_channel

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"regexp"
	"strings"
	"time"
)

var (
	imageRegex *regexp.Regexp
)

func init() {

	// Image regex
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, time.Duration, uint, string,
) {
	// Refer to https://rsshub.app/telegram/channel/nya_sync_dev_test_c

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for telegram channel")

	collectLink =
		strings.ReplaceAll(
			collectLink,
			"{{rsshub_stateless}}",
			config.Config.RSSHubEndpointStateless,
		)

	rawFeed, errCode, err := utils.RSSFeedRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		true,
	)
	if err != nil {
		return false, nil, 0, errCode, err.Error()
	}

	maxFeedIndex := len(rawFeed.Items) - 1

	var feeds []commonTypes.RawFeed
	var minimalInterval time.Duration = work.DropAfter.Sub(work.DropBefore)

	for index, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Title:       item.Title,
				Link:        item.Link,
				GUID:        item.GUID,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			// Process content
			rawContent := item.Description

			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			var normalImages []string
			for _, img := range imgs {
				normalImages = append(normalImages, img[1])
			}

			feed.Media = utils.UploadAllMedia(normalImages)
			for _, media := range feed.Media {
				rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSUri)
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)

			// Calc new interval
			var interv time.Duration
			if index < maxFeedIndex {
				interv = item.PublishedParsed.Sub(*rawFeed.Items[index+1].PublishedParsed)
			} else {
				interv = item.PublishedParsed.Sub(work.DropBefore)
			}
			if interv < 0 {
				interv = -interv
			}
			if interv < minimalInterval {
				minimalInterval = interv
			}
		}
	}

	return true, feeds, minimalInterval, 0, ""

}
