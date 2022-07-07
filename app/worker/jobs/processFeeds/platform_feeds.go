package processFeeds

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/mmcdole/gofeed"
	"math"
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
	posterRegex = regexp.MustCompile("poster=\"(.+?)\"")
	videoRegex = regexp.MustCompile("<source src=\"(.+?)\"")
}

func feedsMedium(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, acceptTime time.Time) {
	// Refer to https://medium.com/feed/@mintable

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	rawFeed, errCode, err := makeRequest(work.CollectLink, true)
	if err != nil {
		handleFailed(work, acceptTime, errCode, err.Error())
		return
	}

	if !strings.Contains(rawFeed.Description, work.VerifyKey) {
		handleFailed(
			work, acceptTime,
			commonConsts.ERROR_CODE_ACCOUNT_NOT_VERIFIED,
			"No identity verify string found on this account",
		)
		return
	}

	var feeds []commonTypes.RawFeed
	var minimalInterval time.Duration = math.MaxInt64

	for index, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feeds = append(feeds, commonTypes.RawFeed{
				Title:       item.Title,
				Link:        item.Link,
				GUID:        item.GUID,
				Categories:  item.Categories,
				Authors:     parseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
				Content:     item.Content,
			})
			if index > 0 {
				interv := rawFeed.Items[index-1].PublishedParsed.Sub(*item.PublishedParsed)
				if interv < 0 {
					interv = -interv
				}
				if interv < minimalInterval {
					minimalInterval = interv
				}
			}
		}
	}

	handleSucceeded(work, acceptTime, feeds, minimalInterval)

}

func feedsTikTok(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, acceptTime time.Time) {
	// Refer to https://github.com/DIYgod/RSSHub/pull/9867

	// Concurrency control
	cccs.Stateful.Request()
	defer cccs.Stateful.Done()

	rawFeed, errCode, err := makeRequest(work.CollectLink, true)
	if err != nil {
		handleFailed(work, acceptTime, errCode, err.Error())
		return
	}

	if !strings.Contains(rawFeed.Description, work.VerifyKey) {
		handleFailed(
			work, acceptTime,
			commonConsts.ERROR_CODE_ACCOUNT_NOT_VERIFIED,
			"No identity verify string found on this account",
		)
		return
	}

	var feeds []commonTypes.RawFeed
	var minimalInterval time.Duration = math.MaxInt64

	for index, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Content:     item.Title, // This should better be content
				PublishedAt: *item.PublishedParsed,
				GUID:        item.GUID,
				Link:        item.Link,
				Authors:     parseAuthors(item.Authors),
			}

			// 2 medias to upload: poster & video
			if posterRegex.MatchString(item.Description) {
				posterUrl := posterRegex.FindStringSubmatch(item.Description)[1]

				// Upload to IPFS
				if ipfsUrl, err := utils.UploadURLToIPFS(posterUrl); err != nil {
					global.Logger.Error("Failed to upload link (", posterUrl, ") onto IPFS: ", err.Error())
				} else {
					posterUrl = ipfsUrl
				}

				// Save to media
				feed.Media = append(feed.Media, posterUrl)
			}
			if videoRegex.MatchString(item.Description) {
				videoUrl := videoRegex.FindStringSubmatch(item.Description)[1]

				// Upload to IPFS
				if ipfsUrl, err := utils.UploadURLToIPFS(videoUrl); err != nil {
					global.Logger.Error("Failed to upload link (", videoUrl, ") onto IPFS: ", err.Error())
				} else {
					videoUrl = ipfsUrl
				}

				// Save to media
				feed.Media = append(feed.Media, videoUrl)
			}

			// Add to results
			feeds = append(feeds, feed)

			// Calc interval
			if index > 0 {
				interv := item.PublishedParsed.Sub(*rawFeed.Items[index-1].PublishedParsed)
				if interv < 0 {
					interv = -interv
				}
				if interv < minimalInterval {
					minimalInterval = interv
				}
			}
		}
	}

	handleSucceeded(work, acceptTime, feeds, minimalInterval)

}

func parseAuthors(authors []*gofeed.Person) []string {
	var parsedAuthors []string
	for _, gofeedAuthor := range authors {
		authorStr := gofeedAuthor.Name
		if gofeedAuthor.Email != "" {
			authorStr += fmt.Sprintf(" <%s>", gofeedAuthor.Email)
		}
		parsedAuthors = append(parsedAuthors, authorStr)
	}

	return parsedAuthors
}
