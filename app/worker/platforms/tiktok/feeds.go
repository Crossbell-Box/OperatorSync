package tiktok

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
	posterRegex *regexp.Regexp
	videoRegex  *regexp.Regexp
)

func init() {

	// TikTok regex
	posterRegex = regexp.MustCompile(`poster="(.+?)"`)
	videoRegex = regexp.MustCompile(`<source src="(.+?)"`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, time.Duration, uint, string,
) {
	// Refer to https://rsshub.app/tiktok/user/@linustech

	// Concurrency control
	cccs.Stateful.Request()
	defer cccs.Stateful.Done()

	global.Logger.Debug("New feeds request for tiktok")

	rawFeed, errCode, err := utils.RSSFeedRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		false,
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
				Title:       item.Title, // This should better be content
				Content:     item.Description,
				PublishedAt: *item.PublishedParsed,
				GUID:        item.GUID,
				Link:        item.Link,
				Authors:     utils.ParseAuthors(item.Authors),
			}

			// 2 medias to upload: poster & video
			if posterRegex.MatchString(feed.Content) {
				rawContent := feed.Content

				posterUrl := posterRegex.FindStringSubmatch(rawContent)[1]

				// Upload to IPFS
				media := commonTypes.Media{
					OriginalURI: posterUrl,
				}
				if media.FileName, media.IPFSUri, media.FileSize, media.ContentType, media.AdditionalProps, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
					global.Logger.Error("Failed to upload poster (", media.OriginalURI, ") onto IPFS: ", err.Error())
				} else {
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSUri)
					feed.Media = append(feed.Media, media)
				}

				feed.Content = rawContent
			}
			if videoRegex.MatchString(feed.Content) {
				rawContent := feed.Content

				videoUrl := videoRegex.FindStringSubmatch(rawContent)[1]

				// Upload to LivePeer, not IPFS
				media := commonTypes.Media{
					OriginalURI: videoUrl,
					FileName:    feed.Title,
				}
				if media.IPFSUri, media.FileSize, media.ContentType, media.AdditionalProps, err = utils.UploadURIToLivePeer(media.OriginalURI, media.FileName); err != nil {
					global.Logger.Error("Failed to upload video (", media.OriginalURI, ") onto LivePeer: ", err.Error())
				} else {
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSUri)
					feed.Media = append(feed.Media, media)
				}

				feed.Content = rawContent
			}

			// Add to results
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
