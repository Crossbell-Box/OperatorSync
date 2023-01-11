package utils

import (
	"encoding/json"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/mmcdole/gofeed"
)

func RSSFeedRequest(url string, withProxy bool) (*gofeed.Feed, uint, error) {
	feedsBody, err := HttpRequest(url, withProxy)
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

func RSSFeedRequestJson(url string, withProxy bool) (*commonTypes.FeedWithExtra, uint, error) {
	feedsBody, err := HttpRequest(url, withProxy)
	if err != nil {
		return nil, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err
	}

	var feed commonTypes.FeedWithExtra

	err = json.Unmarshal(feedsBody, &feed)
	if err != nil {
		return nil, commonConsts.ERROR_CODE_FAILED_TO_PARSE_FEEDS, err
	}

	return &feed, 0, nil
}
