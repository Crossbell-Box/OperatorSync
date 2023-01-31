package tg_channel

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
	imageRegex       *regexp.Regexp
	videoRegex       *regexp.Regexp
	videoPosterRegex *regexp.Regexp
)

func init() {

	// Image regex
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["'].*?/?>`)
	videoRegex = regexp.MustCompile(`<video[^>]+\bsrc=["']([^"']+)["'].*?</video>`)
	videoPosterRegex = regexp.MustCompile(`poster=["']([^"']+)["']`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {
	// Refer to https://rsshub.app/telegram/channel/nya_sync_dev_test_c

	// Concurrency control
	cccs.Stateless.Request()
	defer cccs.Stateless.Done()

	global.Logger.Debug("New feeds request for telegram channel")

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
				//Title:       item.Title,
				Link:        item.URL,
				GUID:        item.ID,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: item.DatePublished,
			}

			// Process content
			rawContent := item.ContentHTML

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

			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			var medias []string
			for _, img := range imgs {
				medias = append(medias, img[1])
			}

			videos := videoRegex.FindAllStringSubmatch(rawContent, -1)
			for _, video := range videos {
				medias = append(medias, video[1])

				// Upload video poster
				poster := videoPosterRegex.FindStringSubmatch(video[0])
				medias = append(medias, poster[1])
			}

			feed.Media = utils.UploadAllMedia(medias)
			for _, media := range feed.Media {
				rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSUri)
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)

		}
	}

	return true, feeds, 0, ""

}
