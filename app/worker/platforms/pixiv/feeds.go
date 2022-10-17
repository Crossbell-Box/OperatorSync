package pixiv

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

func Feeds(cccs *types.ConcurrencyChannels, work *commonTypes.WorkDispatched, collectLink string) (
	bool, []commonTypes.RawFeed, time.Duration, uint, string,
) {

}
