package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"go.uber.org/zap"
	"testing"
)

func TestUploadURLToIPFS(t *testing.T) {

	// Init settings
	config.Config.IPFSEndpoint = "https://ipfs-relay.crossbell.io/upload"

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	// Define variables
	origLink := "https://file.nya.one/misskey/1dfe05b6-32d5-42ff-aa39-7e33aefb84ec.jpg"

	// Test with image
	fileName, ipfsUrl, fileSize, contentType, additionalProps, err := UploadURLToIPFS(origLink, false)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(fileName)
		t.Log(ipfsUrl)         // ipfs://bafkreiftzistch5wftiswc4rkye4zvagbkvdscijhejo43w5bvyjzw7tjm
		t.Log(fileSize)        // 815510
		t.Log(contentType)     // image/jpeg
		t.Log(additionalProps) // {"format":"jpeg","height":"1352","width":"1352"}
	}

}
