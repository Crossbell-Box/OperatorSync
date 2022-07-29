package utils

import commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"

func FeedOnChain(work *commonTypes.OnChainDispatched) (string, error) {
	// TODO : Finish this

	// Step 1: Parse feeds to note metadata

	// Step 2: Upload note metadata to IPFS, get IPFS Uri as note ContentUri

	// Step 3: Upload note to Crossbell Chain with ContentUri

	// Step 4: Return Transaction hash and done!
	return "tx", nil
}
