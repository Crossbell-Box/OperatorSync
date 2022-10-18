package pixiv

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"regexp"
	"strings"
)

var (
	pixivImageRegex *regexp.Regexp
)

func init() {

	// Image regex
	pixivImageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {

	// Refer to https://rsshub.app/pixiv/user/87178177

	// Concurrency control
	cccs.Stateful.Request()
	defer cccs.Stateful.Done()

	global.Logger.Debug("New feeds request for pixiv")

	collectLink =
		strings.ReplaceAll(
			collectLink,
			"{{rsshub_stateful}}",
			config.Config.RSSHubEndpointStateful,
		)

	rawFeed, errCode, err := utils.RSSFeedRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		true,
	)
	if err != nil {
		return false, nil, errCode, err.Error()
	}

	var feeds []commonTypes.RawFeed

	for _, item := range rawFeed.Items {
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

			imgs := pixivImageRegex.FindAllStringSubmatch(rawContent, -1)
			var normalImages []string
			for _, img := range imgs {
				normalImages = append(normalImages, img[1])
				//rawContent = strings.ReplaceAll(rawContent, img[1], "")
			}

			feed.Media = utils.UploadAllMedia(normalImages)

			//feed.Content = rawContent

			feeds = append(feeds, feed)
		}
	}

	return true, feeds, 0, ""

}
