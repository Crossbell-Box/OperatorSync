package chain

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"strconv"
)

func PostNoteForCharacter(characterIdStr string, metadataUri string) (string, error) {
	characterId, err := strconv.ParseInt(characterIdStr, 10, 64)
	if err != nil {
		global.Logger.Errorf("Failed to parse character id with error: %s", err.Error())
		return "", err
	}

	// Prepare instance
	client, auth, err := Prepare()
	if err != nil {
		global.Logger.Errorf("Failed to prepare eth client")
		return "", err
	}

	return "", nil
}
