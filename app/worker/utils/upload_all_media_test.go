package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"go.uber.org/zap"
	"testing"
)

func TestUploadAllMedia(t *testing.T) {

	// Init settings
	config.Config.IPFSEndpoint = "https://ipfs-relay.crossbell.io/upload"

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	// Defile variables
	mediaRegResults := [][]string{{
		"",
		"https://avatars.githubusercontent.com/u/20502130?s=64&amp;v=4",
	}, {
		"",
		"https://avatars.githubusercontent.com/u/20502130?s=64&amp;v=4",
	}, {
		"",
		"https://avatars.githubusercontent.com/u/20502130?s=64&amp;v=4",
	}, {
		"",
		"https://avatars.githubusercontent.com/u/20502130?s=32&amp;v=4",
	}}

	// Run test
	mediaUploadResults := UploadAllMedia(mediaRegResults)
	t.Log("All media files uploaded")
	t.Log(mediaUploadResults)
}
