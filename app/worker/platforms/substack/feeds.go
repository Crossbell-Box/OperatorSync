package substack

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"regexp"
	"strings"
)

var (
	anchorRegex *regexp.Regexp
	imageRegex  *regexp.Regexp
)

func init() {

	// Image regex
	anchorRegex = regexp.MustCompile(`<a[\s\S]+?</a>`)
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)

}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {
	// Refer to https://lc499.substack.com/feed

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for substack")

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
				Description: item.Description,
				Link:        item.Link,
				GUID:        item.GUID,
				Categories:  item.Categories,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			// Find enclosure
			for _, e := range item.Enclosures {
				if strings.HasPrefix(e.Type, "image/") {
					uploadedImg := utils.UploadAllMedia([]string{e.URL})
					if len(uploadedImg) > 0 {
						feed.Image = uploadedImg[0].IPFSUri
						break
					}
				}
			}

			// Process content
			rawContent := item.Content
			anchors := anchorRegex.FindAllString(rawContent, -1)
			var images []string
			img2a := make(map[string]string)
			for _, a := range anchors {
				imgs := imageRegex.FindAllStringSubmatch(a, -1)
				for _, img := range imgs {
					images = append(images, img[1])
					img2a[img[1]] = a
				}
			}

			feed.Media = utils.UploadAllMedia(images)
			for _, media := range feed.Media {
				if a, ok := img2a[media.OriginalURI]; ok {
					rawContent = strings.ReplaceAll(rawContent, a, media.IPFSUri)
				} else {
					// Unable to find original anchor, so just replace the link
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSUri)
				}
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)

		}
	}

	return true, feeds, 0, ""

}
