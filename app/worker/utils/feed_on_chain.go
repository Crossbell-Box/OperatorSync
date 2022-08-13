package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/chain"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"strconv"
)

func FeedOnChain(work *commonTypes.OnChainRequest) (string, string, error) {
	// Step 1: Parse feeds to note metadata
	metadata := types.NoteMetadata{
		Type: "note",
		Authors: append(
			work.Authors,
			fmt.Sprintf("csb://account:%s@%s", work.Username, work.Platform),
		),
		Title:   work.Title,
		Content: work.Content,
		Tags:    work.Categories,
		Sources: []string{
			"OperatorSync",
			commonConsts.SUPPORTED_PLATFORM[work.Platform].Name,
		},
		ContentWarning: work.ContentWarning,
		DatePublished:  work.PublishedAt.Format("2022-01-01T00:00:00Z"),
	}

	if ValidateUri(work.Link) {
		metadata.ExternalUrls = []string{work.Link}
	}

	for _, media := range work.Media {
		// Append basic info
		attachment := types.NoteAttachment{
			Name:    media.FileName,
			Address: media.IPFSUri,
			//Content:  "",
			MimeType: media.ContentType,
			FileSize: media.FileSize,
			Alt:      media.FileName, // Just use filename for now
		}

		// Append additional props
		var additionalProps map[string]string
		err := json.Unmarshal([]byte(media.AdditionalProps), &additionalProps)
		if err != nil {
			global.Logger.Errorf("Failed to parse additional props for media #%s with error: %s", media.IPFSUri, err.Error())
		} else {
			// If has "width" and "height"
			if mediaWidthStr, ok := additionalProps["width"]; ok {
				mediaWidth, err := strconv.Atoi(mediaWidthStr)
				if err != nil {
					global.Logger.Errorf("Failed to parse width of media #%s with error: %s", media.IPFSUri, err.Error())
				} else {
					attachment.Width = uint(mediaWidth)
				}
			}
			if mediaHeightStr, ok := additionalProps["height"]; ok {
				mediaHeight, err := strconv.Atoi(mediaHeightStr)
				if err != nil {
					global.Logger.Errorf("Failed to parse height of media #%s with error: %s", media.IPFSUri, err.Error())
				} else {
					attachment.Width = uint(mediaHeight)
				}
			}
		}

		// Append to array
		metadata.Attachments = append(metadata.Attachments, attachment)
	}

	// Step 2: Upload note metadata to IPFS, get IPFS Uri as note ContentUri
	metaBytes, err := json.Marshal(&metadata)
	if err != nil {
		global.Logger.Errorf("Failed to parse metadata to json with error: %s", err.Error())
		return "", "", err
	}
	ipfsUri, _, err := UploadBytesToIPFS(metaBytes, "metadata.json")
	if err != nil {
		global.Logger.Errorf("Failed to upload metadata to IPFS with error: %s", err.Error())
		return "", "", err
	}

	// Step 3: Upload note to Crossbell Chain with ContentUri
	tx, err := chain.PostNoteForCharacter(work.CrossbellCharacterID, ipfsUri)
	if err != nil {
		global.Logger.Errorf("Failed to post note to Crossbell chain for character #%s with error: %s", work.CrossbellCharacterID, err.Error())
		return ipfsUri, tx, err // Transaction might be invalid
	}

	// Step 4: Return Transaction hash and done!
	return ipfsUri, tx, nil
}
