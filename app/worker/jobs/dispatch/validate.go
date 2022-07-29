package dispatch

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"net/http"
)

type crossbellIndexerCharacterRes struct {
	//CharacterID uint `json:"characterId"`
	Handle string `json:"handle"`
	// Ignore others
}

func ValidateAccounts(m *nats.Msg) {
	global.Logger.Debug("New validate request received: ", string(m.Data))

	var validateReq commonTypes.ValidateRequest

	if err := json.Unmarshal(m.Data, &validateReq); err != nil {
		global.Logger.Error("Failed to parse validate request.", err.Error())
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_FAILED_TO_PARSE_JSON, "Failed to parse validate request")
		return
	}

	// With character ID https://indexer.crossbell.io/v1/characters/19
	// With handle https://indexer.crossbell.io/v1/handles/candinya/character

	reqUrl := fmt.Sprintf("https://indexer.crossbell.io/v1/characters/%s", validateReq.CrossbellCharacterID)
	rawRes, err := (&http.Client{}).Get(reqUrl)
	if err != nil {
		global.Logger.Error("Failed to get character info for ", validateReq.CrossbellCharacterID, " : ", err.Error())
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, "Failed to get character info")
		return
	}

	var crossbellIndexerResponse crossbellIndexerCharacterRes
	if err = json.NewDecoder(rawRes.Body).Decode(&crossbellIndexerResponse); err != nil {
		global.Logger.Error("Failed to parse response data: ", err.Error())
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_FAILED_TO_PARSE_JSON, "Failed to parse response data")
		return
	}

	validateString := fmt.Sprintf("Crossbell@%s#%s", crossbellIndexerResponse.Handle, validateReq.CrossbellCharacterID)

	switch validateReq.Platform {
	case "medium":
		medium.Account(m.Reply, validateReq.Username, validateString)
	case "tiktok":
		tiktok.Account(m.Reply, validateReq.Username, validateString)
	default:
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
	}

}
