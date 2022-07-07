package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"testing"
)

func TestUploadURLToIPFS(t *testing.T) {

	// Init settings
	config.Config.IPFSEndpoint = "https://ipfs-relay.crossbell.io/upload"

	// Define variables
	origLink := "https://file.nya.one/misskey/1dfe05b6-32d5-42ff-aa39-7e33aefb84ec.jpg"

	// Test with image
	ipfsUrl, err := UploadURLToIPFS(origLink)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(ipfsUrl) // ipfs://bafkreiftzistch5wftiswc4rkye4zvagbkvdscijhejo43w5bvyjzw7tjm
	}

}
