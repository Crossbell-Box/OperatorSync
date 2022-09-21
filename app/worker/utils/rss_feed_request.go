package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/mmcdole/gofeed"
)

func RSSFeedRequest(url string, withProxy bool) (*gofeed.Feed, uint, error) {
	feedsBody, err := HttpRequest(url, withProxy)
	if err != nil {
		return nil, consts.ERROR_CODE_HTTP_REQUEST_FAILED, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(feedsBody))
	if err != nil {
		return nil, consts.ERROR_CODE_FAILED_TO_PARSE_FEEDS, err
	}

	global.Logger.Debug("Parsed feeds: ", feed)

	return feed, 0, nil
}
