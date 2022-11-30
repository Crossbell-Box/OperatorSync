package jobs

import (
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strings"
	"testing"
)

func TestSplitConnectedAccounts(t *testing.T) {
	accountPlatformSplits := strings.Split("lc499@thenomads.social@mastodon", "@")
	t.Log(commonTypes.Account{
		Platform: accountPlatformSplits[len(accountPlatformSplits)-1],
		Username: strings.Join(accountPlatformSplits[0:len(accountPlatformSplits)-1], "@"), // For fediverse
	})
}
