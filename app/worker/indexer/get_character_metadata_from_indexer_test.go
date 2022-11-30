package indexer

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"go.uber.org/zap"
	"testing"
)

func TestGetCharacterMetadataFromIndexer(t *testing.T) {

	// Init settings
	config.Config.CrossbellIndexer = "https://indexer.crossbell.io"

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	// Defile variables
	characterId1 := "451"
	characterId2 := "19"

	// Run test
	t.Log(GetCharacterMetadataFromIndexer(characterId1))
	t.Log(GetCharacterMetadataFromIndexer(characterId2))
}
