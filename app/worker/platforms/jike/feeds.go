package jike

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"regexp"
	"strings"
)

var (
	imageRegex *regexp.Regexp
)

func init() {
	// Image regex
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {
	// Refer to https://rsshub.app/jike/user/3EE02BC9-C5B3-4209-8750-4ED1EE0F67BB

	// Concurrency control
	cccs.Stateless.Request()
	defer cccs.Stateless.Done()

	global.Logger.Debug("New feeds request for jike")

	collectLink =
		strings.ReplaceAll(
			collectLink,
			"{{rsshub_stateless}}",
			config.Config.RSSHubEndpointStateless,
		)

	rawFeed, errCode, err := utils.RSSFeedRequestJson(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		true,
	)
	if err != nil {
		return false, nil, errCode, err.Error()
	}

	var feeds []commonTypes.RawFeed

	for _, item := range rawFeed.Items {
		if item.DatePublished.After(work.DropBefore) && item.DatePublished.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Link:        item.URL,
				GUID:        item.ID,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: item.DatePublished,
			}

			// Process content
			rawContent := item.ContentHTML // Different from common XML feeds

			skipThis := false

			// Check extra links
			for _, link := range item.Extra.Links {
				switch link.Type {
				case "quote":
					// Set ForURI
					feed.ForURI = link.URL
					// Remove inner quote content
					rawContent = strings.Replace(rawContent, link.ContentHTML, "", 1)

				case "reply":
					// Set ForURI
					feed.ForURI = link.URL

				case "repost":
					// No need to post, skip this
					skipThis = true

				default:
					// Unknown link, just log it
					global.Logger.Warnf("Unknown extra link type detected: %v", link)
				}
			}

			if skipThis {
				continue
			}

			var medias []string
			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			for _, img := range imgs {
				medias = append(medias, img[1])
				rawContent = strings.ReplaceAll(rawContent, img[0], "")
			}

			// Upload media with order
			for _, mediaUri := range medias {
				media, err := utils.UploadOneMedia(mediaUri)
				if err != nil {
					// Fail to upload, oops
					return false, nil, commonConsts.ERROR_CODE_FAILED_TO_UPLOAD, err.Error()
				} else {
					feed.Media = append(feed.Media, *media)
				}
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)

		}
	}

	return true, feeds, 0, ""

}
