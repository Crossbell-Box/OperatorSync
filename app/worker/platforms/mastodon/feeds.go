package mastodon

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	commonUtils "github.com/Crossbell-Box/OperatorSync/common/utils"
	"regexp"
	"strings"
)

var (
	customEmojisRegex *regexp.Regexp
)

func init() {

	// Medium regex
	customEmojisRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {
	// Refer to https://eihei.net/@candinya.rss

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for mastodon")

	username, instance, err := commonUtils.SplitFediverseUsernameInstance(work.Username)
	if err != nil {
		return false, nil, commonConsts.ERROR_CODE_INVALID_FORMAT, err.Error()
	}

	collectLink = strings.Replace(collectLink, "{{instance}}", instance, 1)
	collectLink = strings.Replace(collectLink, "{{username}}", username, 1)

	rawFeed, errCode, err := utils.RSSFeedRequest(
		collectLink,
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
				PublishedAt: *item.PublishedParsed,
			}

			// Process content: upload custom emojis
			rawContent := item.Description

			var customEmojis []string
			ces := customEmojisRegex.FindAllStringSubmatch(rawContent, -1)
			for _, ce := range ces {
				customEmojis = append(customEmojis, ce[1])
			}

			ceUploadeds := utils.UploadAllMedia(customEmojis)
			for _, ceu := range ceUploadeds {
				rawContent = strings.ReplaceAll(rawContent, ceu.OriginalURI, ceu.IPFSUri)
			}

			feed.Content = rawContent

			// Process medias
			if attachedMedias, ok := item.Extensions["media"]["content"]; ok {
				// Upload media with order
				for _, aMedia := range attachedMedias {
					media, err := utils.UploadOneMedia(aMedia.Attrs["url"])
					if err != nil {
						// Fail to upload, oops
						return false, nil, commonConsts.ERROR_CODE_FAILED_TO_UPLOAD, err.Error()
					} else {
						feed.Media = append(feed.Media, *media)
					}
				}
			}

			feeds = append(feeds, feed)

		}
	}

	return true, feeds, 0, ""
}
