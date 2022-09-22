package medium

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"net/url"
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

func removeTrackingQuery(rawUri string) string {
	// Process link: remove tracking query

	rawLinkUri, err := url.Parse(rawUri)
	if err != nil {
		return rawUri
	}
	// Only process when succeeded to parse URI
	rawUriQuery := rawLinkUri.Query()
	if rawUriQuery.Has("source") {
		rawUriQuery.Del("source")
		// Delete `source=rss-xxxxxxxxxxxx------x`
	}
	rawLinkUri.RawQuery = rawUriQuery.Encode()
	return rawLinkUri.String()
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, time.Duration, uint, string,
) {
	// Refer to https://medium.com/feed/@nya_9949

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for medium")

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
				Link:        removeTrackingQuery(item.Link),
				GUID:        item.GUID,
				Categories:  item.Categories,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			// Process content: remove tracking image
			rawContent := item.Content

			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			// Remove tracker
			var normalImages []string
			for _, img := range imgs {
				if strings.Contains(img[1], "https://medium.com/_/stat") {
					rawContent = strings.ReplaceAll(
						rawContent,
						fmt.Sprintf("<img src=\"%s\" width=\"1\" height=\"1\" alt=\"\">", img[1]),
						"",
					)
				} else {
					normalImages = append(normalImages, img[1])
				}
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
