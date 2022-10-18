package tiktok

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"regexp"
	"strings"
	"time"
)

var (
	//posterRegex *regexp.Regexp
	videoRegex *regexp.Regexp
)

func init() {

	// TikTok regex
	//posterRegex = regexp.MustCompile(`poster="(.+?)"`)
	videoRegex = regexp.MustCompile(`<source src="(.+?)"`)
}

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, time.Duration, uint, string,
) {
	// Refer to https://rsshub.app/tiktok/user/@linustech

	// Concurrency control
	cccs.Stateless.Request()
	defer cccs.Stateless.Done()

	global.Logger.Debug("New feeds request for tiktok")

	collectLink =
		strings.ReplaceAll(
			collectLink,
			"{{rsshub_stateless}}",
			config.Config.RSSHubEndpointStateless,
		)

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
				Title: item.Title,
				//Content:     item.Description,
				PublishedAt: *item.PublishedParsed,
				GUID:        item.GUID,
				Link:        item.Link,
				Authors:     utils.ParseAuthors(item.Authors),
			}

			// 2 medias to upload: poster & video
			rawContent := item.Description
			//if posterRegex.MatchString(rawContent) {
			//
			//	posterUrl := posterRegex.FindStringSubmatch(rawContent)[1]
			//
			//	// Upload to IPFS
			//	media := commonTypes.Media{
			//		OriginalURI: posterUrl,
			//	}
			//	if media.FileName, media.IPFSUri, media.FileSize, media.ContentType, media.AdditionalProps, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
			//		global.Logger.Error("Failed to upload poster (", media.OriginalURI, ") onto IPFS: ", err.Error())
			//		// Still acceptable
			//	} else {
			//		feed.Media = append(feed.Media, media)
			//	}
			//}
			if videoRegex.MatchString(rawContent) {

				videoUrl := videoRegex.FindStringSubmatch(rawContent)[1]

				// Upload to IPFS
				media := commonTypes.Media{
					OriginalURI: videoUrl,
				}
				if media.FileName, media.IPFSUri, media.FileSize, media.ContentType, media.AdditionalProps, err = utils.UploadURLToIPFS(media.OriginalURI, true); err != nil {
					global.Logger.Error("Failed to upload video (", media.OriginalURI, ") onto IPFS: ", err.Error())
					// Unacceptable
					return false, nil, 0, commonConsts.ERROR_CODE_FAILED_TO_UPLOAD, err.Error()
				} else {
					feed.Media = append(feed.Media, media)
				}
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
