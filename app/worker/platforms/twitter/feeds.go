package twitter

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
	imageRegex *regexp.Regexp
	videoRegex *regexp.Regexp
)

func init() {

	// Image regex
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)
	videoRegex = regexp.MustCompile(`<video[^>]+\bsrc=["']([^"']+)["'].*?</video>`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {
	// Refer to https://rsshub.app/twitter/user/lc499

	// Concurrency control
	cccs.Stateless.Request()
	defer cccs.Stateless.Done()

	global.Logger.Debug("New feeds request for twitter")

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
		return false, nil, errCode, err.Error()
	}

	var feeds []commonTypes.RawFeed

	for _, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Link:        item.Link,
				GUID:        item.GUID,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			// Process content
			rawContent := item.Description

			var medias []string
			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			for _, img := range imgs {
				medias = append(medias, img[1])
				rawContent = strings.ReplaceAll(rawContent, img[0], "")
			}

			videos := videoRegex.FindAllStringSubmatch(rawContent, -1)
			for _, video := range videos {
				medias = append(medias, video[1])
				rawContent = strings.ReplaceAll(rawContent, video[0], "")
			}

			feed.Media = utils.UploadAllMedia(medias)

			feed.Content = rawContent

			feeds = append(feeds, feed)

		}
	}

	return true, feeds, 0, ""

}
