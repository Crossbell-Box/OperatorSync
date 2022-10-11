package pinterest

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"regexp"
	"strings"
	"time"
)

var (
	anchorRegex   *regexp.Regexp
	imageRegex    *regexp.Regexp
	leadingSpaces *regexp.Regexp
)

func init() {

	// Image regex
	anchorRegex = regexp.MustCompile(`<a[\s\S]+?</a>`)
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)
	leadingSpaces = regexp.MustCompile(`^[\s\n\r]+`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, time.Duration, uint, string,
) {
	// Refer to https://www.pinterest.com/ncandy0418/feed.rss

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for pinterest")

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

			var originalSizeImgs []string

			rawContent := item.Description
			anchors := anchorRegex.FindAllString(rawContent, -1)
			for _, a := range anchors {
				imgs := imageRegex.FindAllStringSubmatch(a, -1)
				for _, img := range imgs {
					originalsUri := strings.Replace(img[1], "/236x/", "/originals/", 1)
					originalSizeImgs = append(originalSizeImgs, originalsUri)
					rawContent = strings.Replace(rawContent, a, "", 1)
				}
			}

			feed.Media = utils.UploadAllMedia(originalSizeImgs)

			// Remove leading spaces
			feed.Content = leadingSpaces.ReplaceAllString(rawContent, "")

			feeds = append(feeds, feed)

			// Calc interval
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
