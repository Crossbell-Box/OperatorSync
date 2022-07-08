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
	imageRegex  *regexp.Regexp
	posterRegex *regexp.Regexp
	videoRegex  *regexp.Regexp
)

func init() {

	// Medium regex
	imageRegex = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

	// TikTok regex
	posterRegex = regexp.MustCompile(`poster="(.+?)"`)
	videoRegex = regexp.MustCompile(`<source src="(.+?)"`)
}

func feedsMedium(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, acceptTime time.Time, collectLink string) {
	// Refer to https://medium.com/feed/@nya_9949

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for medium")

	rawFeed, errCode, err := makeRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		true,
	)
	if err != nil {
		handleFailed(work, acceptTime, errCode, err.Error())
		return
	}

	// Medium feed won't show user's bio
	//if !strings.Contains(strings.ToLower(rawFeed.Description), strings.ToLower(work.VerifyKey)) {
	//	handleFailed(
	//		work, acceptTime,
	//		commonConsts.ERROR_CODE_ACCOUNT_NOT_VERIFIED,
	//		"No identity verify string found on this account",
	//	)
	//	return
	//}

	var feeds []commonTypes.RawFeed
	var minimalInterval time.Duration = math.MaxInt64

	for index, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Title:       item.Title,
				Link:        item.Link,
				GUID:        item.GUID,
				Categories:  item.Categories,
				Authors:     parseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
			}

			rawContent := item.Content

			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			for _, img := range imgs {
				media := commonTypes.Media{
					OriginalURI: img[1],
				}
				if media.IPFSURI, media.FileSize, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
					global.Logger.Error("Failed to upload link (", media.OriginalURI, ") onto IPFS: ", err.Error())
				} else {
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSURI)
					feed.Media = append(feed.Media, media)
				}
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)
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

func feedsGitHub(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, acceptTime time.Time, collectLink string) {
	// Refer to https://github.com/candinya.atom

	// Concurrency control
	cccs.Direct.Request()
	defer cccs.Direct.Done()

	global.Logger.Debug("New feeds request for github")

	rawFeed, errCode, err := makeRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		true,
	)
	if err != nil {
		handleFailed(work, acceptTime, errCode, err.Error())
		return
	}

	if !rawFeed.UpdatedParsed.After(work.DropBefore) {
		// Not updated after last update, drop all
		handleSucceeded(work, acceptTime, nil, 0)
		return
	}

	// GitHub feed won't show user's bio
	//if !strings.Contains(strings.ToLower(rawFeed.Description), strings.ToLower(work.VerifyKey)) {
	//	handleFailed(
	//		work, acceptTime,
	//		commonConsts.ERROR_CODE_ACCOUNT_NOT_VERIFIED,
	//		"No identity verify string found on this account",
	//	)
	//	return
	//}

	var feeds []commonTypes.RawFeed
	var minimalInterval time.Duration = math.MaxInt64

	for index, item := range rawFeed.Items {
		if item.PublishedParsed.After(work.DropBefore) && item.PublishedParsed.Before(work.DropAfter) {
			feed := commonTypes.RawFeed{
				Title:       item.Title,
				Link:        item.Link,
				GUID:        item.GUID,
				Authors:     parseAuthors(item.Authors),
				PublishedAt: *item.PublishedParsed,
				UpdatedAt:   *item.UpdatedParsed,
			}

			rawContent := item.Content

			imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
			for _, img := range imgs {
				media := commonTypes.Media{
					OriginalURI: img[1],
				}
				if media.IPFSURI, media.FileSize, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
					global.Logger.Error("Failed to upload link (", media.OriginalURI, ") onto IPFS: ", err.Error())
				} else {
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSURI)
					feed.Media = append(feed.Media, media)
				}
			}

			feed.Content = rawContent

			feeds = append(feeds, feed)
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

func feedsTikTok(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, acceptTime time.Time, collectLink string) {
	// Refer to https://rsshub.app/tiktok/user/@linustech

	// Concurrency control
	cccs.Stateful.Request()
	defer cccs.Stateful.Done()

	global.Logger.Debug("New feeds request for tiktok")

	rawFeed, errCode, err := makeRequest(
		strings.ReplaceAll(collectLink, "{{username}}", work.Username),
		false,
	)
	if err != nil {
		handleFailed(work, acceptTime, errCode, err.Error())
		return
	}

	if !strings.Contains(strings.ToLower(rawFeed.Description), strings.ToLower(work.VerifyKey)) {
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
				Title:       item.Title, // This should better be content
				Content:     item.Description,
				PublishedAt: *item.PublishedParsed,
				GUID:        item.GUID,
				Link:        item.Link,
				Authors:     parseAuthors(item.Authors),
			}

			// 2 medias to upload: poster & video
			if posterRegex.MatchString(feed.Content) {
				rawContent := feed.Content

				posterUrl := posterRegex.FindStringSubmatch(rawContent)[1]

				// Upload to IPFS
				media := commonTypes.Media{
					OriginalURI: posterUrl,
				}
				if media.IPFSURI, media.FileSize, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
					global.Logger.Error("Failed to upload link (", media.OriginalURI, ") onto IPFS: ", err.Error())
				} else {
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSURI)
					feed.Media = append(feed.Media, media)
				}

				feed.Content = rawContent
			}
			if videoRegex.MatchString(feed.Content) {
				rawContent := feed.Content

				videoUrl := videoRegex.FindStringSubmatch(rawContent)[1]

				// Upload to IPFS
				media := commonTypes.Media{
					OriginalURI: videoUrl,
				}
				if media.IPFSURI, media.FileSize, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
					global.Logger.Error("Failed to upload link (", media.OriginalURI, ") onto IPFS: ", err.Error())
				} else {
					rawContent = strings.ReplaceAll(rawContent, media.OriginalURI, media.IPFSURI)
					feed.Media = append(feed.Media, media)
				}

				feed.Content = rawContent
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
