package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/common/types"
	"sync"
)

func UploadAllMedia(mediaUris []string) []types.Media {

	// Collect all unique media URIs
	mediaUriSet := make(map[string]struct{})
	for _, rawMediaUri := range mediaUris {
		if _, ok := mediaUriSet[rawMediaUri]; !ok {
			mediaUriSet[rawMediaUri] = struct{}{}
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
			if media.FileName, media.IPFSUri, media.FileSize, media.ContentType, media.AdditionalProps, err = UploadURLToIPFS(media.OriginalURI); err != nil {
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
