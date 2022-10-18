package utils

import (
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func CalcNewInterval(work *commonTypes.WorkDispatched, feeds []commonTypes.RawFeed) time.Duration {

	var minimalInterval time.Duration = work.DropAfter.Sub(work.DropBefore)
	maxFeedIndex := len(feeds) - 1
	for index, item := range feeds {

		// Calc new interval
		var interv time.Duration
		if index < maxFeedIndex {
			interv = item.PublishedAt.Sub(feeds[index+1].PublishedAt)
		} else {
			interv = item.PublishedAt.Sub(work.DropBefore)
		}
		if interv < 0 {
			interv = -interv
		}
		if interv < minimalInterval {
			minimalInterval = interv
		}
	}

	return minimalInterval
}
