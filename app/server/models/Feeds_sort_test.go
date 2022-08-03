package models

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"sort"
	"testing"
	"time"
)

func TestSortFeeds(t *testing.T) {
	basetime := time.Now()

	feeds := FeedsArray{
		{
			Feed: types.Feed{
				RawFeed: commonTypes.RawFeed{
					Title:       "Demo 3",
					PublishedAt: basetime.Add(time.Hour),
				},
			},
		},
		{
			Feed: types.Feed{
				RawFeed: commonTypes.RawFeed{
					Title:       "Demo 2",
					PublishedAt: basetime.Add(time.Hour),
				},
			},
		},
		{
			Feed: types.Feed{
				RawFeed: commonTypes.RawFeed{
					Title:       "Demo 1",
					PublishedAt: basetime.Add(-time.Hour),
				},
			},
		},
	}

	sort.Sort(feeds)

	if feeds[0].PublishedAt != basetime.Add(-time.Hour) || feeds[0].Title != "Demo 1" {
		t.Fail()
	}

}
