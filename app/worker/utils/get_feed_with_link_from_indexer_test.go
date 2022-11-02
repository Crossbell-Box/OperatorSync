package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"go.uber.org/zap"
	"testing"
)

func TestGetFeedWithLinkFromIndexer(t *testing.T) {

	// Init settings
	config.Config.CrossbellIndexer = "https://indexer.crossbell.io"

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	// Defile variables
	link := "https://twitter.com/NyaRSS3/status/1534713583270998018"

	// Run test
	t.Log(GetFeedWithLinkFromIndexer(link))
}
