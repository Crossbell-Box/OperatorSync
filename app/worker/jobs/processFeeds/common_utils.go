package processFeeds

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/mmcdole/gofeed"
	"sync"
	"time"
)

func makeRequest(url string, withProxy bool) (*gofeed.Feed, uint, error) {
	feedsBody, err := utils.HttpRequest(url, withProxy)
	if err != nil {
		return nil, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(feedsBody))
	if err != nil {
		return nil, commonConsts.ERROR_CODE_FAILED_TO_PARSE_FEEDS, err
	}

	return feed, 0, nil
}

func handleSucceeded(workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, rawFeeds []commonTypes.RawFeed, newInterval time.Duration) {

	global.Logger.Debug("Work succeeded: ", workDispatched)

	if len(rawFeeds) == 0 {
		// Actually nothing
		global.Logger.Debug("... but nothing found")
		handleFailed(workDispatched, acceptTime, commonConsts.ERROR_CODE_NOTHING_FOUND, "Nothing found")
		return
	}

	succeededWork := commonTypes.WorkSucceeded{
		WorkDispatched: *workDispatched,
		AcceptedAt:     acceptTime,
		SucceededAt:    time.Now(),
		Feeds:          rawFeeds,
		NewInterval:    newInterval,
	}
	if succeededWorkBytes, err := json.Marshal(&succeededWork); err != nil {
		global.Logger.Error("Failed to marshall succeeded work: ", succeededWork)
	} else {
		err = global.MQ.Publish(commonConsts.MQSETTINGS_SucceededChannelName, succeededWorkBytes)
		if err != nil {
			global.Logger.Error("Failed to report succeeded work with error: ", err.Error())
		} else {
			global.Logger.Debug("And result pushed to succeeded channel.")
		}
	}
}

func handleFailed(workDispatched *commonTypes.WorkDispatched, acceptTime time.Time, errCode uint, errMsg string) {
	global.Logger.Warn("Work failed: ", workDispatched, " with code: ", errCode, " , reason: ", errMsg)

	failedWork := commonTypes.WorkFailed{
		WorkDispatched: *workDispatched,
		AcceptAt:       acceptTime,
		ErrorAt:        time.Now(),
		ErrorCode:      errCode,
		ErrorReason:    errMsg,
	}
	if failedWorkBytes, err := json.Marshal(&failedWork); err != nil {
		global.Logger.Error("Failed to marshall failed work: ", failedWork)
	} else {
		err = global.MQ.Publish(commonConsts.MQSETTINGS_FailedChannelName, failedWorkBytes)
		if err != nil {
			global.Logger.Error("Failed to report failed work with error: ", err.Error())
		}
	}
}

func uploadAllMedia(regResult [][]string) []commonTypes.Media {

	// Collect all unique media URIs
	mediaUriSet := make(map[string]struct{})
	for _, rawMediaUriRegRes := range regResult {
		if _, ok := mediaUriSet[rawMediaUriRegRes[1]]; !ok {
			mediaUriSet[rawMediaUriRegRes[1]] = struct{}{}
		}
	}

	// Upload them all
	var ipfsUploadWg sync.WaitGroup
	ipfsUploadResultChannel := make(chan commonTypes.Media, len(mediaUriSet))
	for uri := range mediaUriSet {
		uri := uri
		ipfsUploadWg.Add(1)
		go func() {
			media := commonTypes.Media{
				OriginalURI: uri,
			}
			var err error
			if media.IPFSURI, media.FileSize, err = utils.UploadURLToIPFS(media.OriginalURI); err != nil {
				global.Logger.Error("Failed to upload link (", media.OriginalURI, ") onto IPFS: ", err.Error())
			} else {
				ipfsUploadResultChannel <- media
			}
			ipfsUploadWg.Done()
		}()
	}

	// Collect results
	ipfsUploadWg.Wait()
	close(ipfsUploadResultChannel)
	var medias []commonTypes.Media
	for media := range ipfsUploadResultChannel {
		medias = append(medias, media)
	}

	return medias

}
