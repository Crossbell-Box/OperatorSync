package utils

import (
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func FeedOnChain(work *commonTypes.OnChainDispatched) (string, string, error) {
	// Step 1: Parse feeds to note metadata
	metadata := types.NoteMetadata{
		Type:    "note",
		Title:   work.Title,
		Content: work.Content,
		Tags:    work.Categories,
		//Attachments: nil,
		Sources: []string{
			"OperatorSync",
			work.Platform,
		},
	}

	for _, media := range work.Media {
		metadata.Attachments = append(metadata.Attachments, types.NoteAttachment{
			Name:    media.FileName,
			Address: media.IPFSUri,
			//Content:  "",
			MimeType: media.ContentType,
			FileSize: media.FileSize,
		})
	}

	// Step 2: Upload note metadata to IPFS, get IPFS Uri as note ContentUri
	metaBytes, err := json.Marshal(&metadata)
	if err != nil {
		global.Logger.Error("Failed to parse metadata to json with error: %s", err.Error())
		return "", "", err
	}
	ipfsUri, _, err := UploadBytesToIPFS(metaBytes, "metadata.json")
	if err != nil {
		global.Logger.Error("Failed to upload metadata to IPFS with error: %s", err.Error())
		return "", "", err
	}

	// Step 3: Upload note to Crossbell Chain with ContentUri
	// TODO : Upload note to Crossbell

	// Step 4: Return Transaction hash and done!
	return ipfsUri, "tx", nil
}
