package processFeeds

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"strings"
	"time"
)

func ProcessFeeds(platformID string, cccs *types.ConcurrencyChannels) func(m *nats.Msg) {
	return func(m *nats.Msg) {

		var workDispatched commonTypes.WorkDispatched

		if err := json.Unmarshal(m.Data, &workDispatched); err != nil {
			global.Logger.Error("Failed to parse dispatched work data.")
			// Even unable to report to failed works
			return
		}

		acceptTime := time.Now()

		// Parse settings
		workDispatched.CollectLink =
			strings.ReplaceAll(
				strings.ReplaceAll(
					workDispatched.CollectLink,
					"{{rsshub_stateful}}",
					config.Config.RSSHubEndpointStateful,
				),
				"{{rsshub_stateless}}",
				config.Config.RSSHubEndpointStateless,
			)

		switch platformID {
		case "medium":
			feedsMedium(cccs, &workDispatched, acceptTime)
		case "tiktok":
			feedsTikTok(cccs, &workDispatched, acceptTime)
		default:
			// Unable to handle
			handleFailed(&workDispatched, acceptTime, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
		}
	}
}
