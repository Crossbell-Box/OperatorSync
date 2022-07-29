package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/common/types"
	"sync"
)

func UploadAllMedia(regResult [][]string) []types.Media {

	// Collect all unique media URIs
	mediaUriSet := make(map[string]struct{})
	for _, rawMediaUriRegRes := range regResult {
		if _, ok := mediaUriSet[rawMediaUriRegRes[1]]; !ok {
			mediaUriSet[rawMediaUriRegRes[1]] = struct{}{}
		}
	}

	// Upload them all
	var ipfsUploadWg sync.WaitGroup
	ipfsUploadResultChannel := make(chan types.Media, len(mediaUriSet))
	for uri := range mediaUriSet {
		uri := uri
		ipfsUploadWg.Add(1)
		go func() {
			media := types.Media{
				OriginalURI: uri,
			}
			var err error
			if media.IPFSUri, media.FileSize, media.ContentType, err = UploadURLToIPFS(media.OriginalURI); err != nil {
				global.Logger.Error("Failed to upload link (", media.OriginalURI, ") onto IPFS: ", err.Error())
			} else {
				ipfsUploadResultChannel <- media
			}
			ipfsUploadWg.Done()
		}()
	}

	// Collect results
	ipfsUploadWg.Wait()
	close(ipfsUploadResultChannel)
	var medias []types.Media
	for media := range ipfsUploadResultChannel {
		medias = append(medias, media)
	}

	return medias

}
