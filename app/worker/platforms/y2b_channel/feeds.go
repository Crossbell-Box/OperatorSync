package y2b_channel

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strings"
)

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, uint, string,
) {

	// Refer to https://www.youtube.com/feeds/videos.xml?channel_id=UCs_f32UpeFk2CM8hzF7Q61Q

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for YouTube Channel")

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
				Content:     item.Extensions["media"]["group"][0].Children["description"][0].Value,
				Link:        item.Link,
				GUID:        item.GUID,
				Authors:     utils.ParseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			// Process content: download video
			var targetVideo commonTypes.Media

			targetVideo.IPFSUri, targetVideo.FileSize, err = utils.UploadVideoToIPFS(feed.Link)
			if err != nil {
				global.Logger.Errorf("Failed to upload video (%s) to IPFS with error: %s", feed.Link, err.Error())
				// Unacceptable
				return false, nil, commonConsts.ERROR_CODE_FAILED_TO_UPLOAD, err.Error()
			} else {
				feed.Media = append(feed.Media, targetVideo)
			}

			// Add to results
			feeds = append(feeds, feed)

		}
	}

	return true, feeds, 0, ""

}
