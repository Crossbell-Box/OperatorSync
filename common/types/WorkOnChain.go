package types

type OnChainDispatched struct {
	FeedID               uint   `json:"feed_id"` // Feed ID in main database
	CrossbellCharacterID string `json:"crossbell_character_id"`
	Platform             string `json:"platform"`
	RawFeed
}

type OnChainRespond struct {
	FeedID      uint   `json:"feed_id"` // Feed ID in main database
	IPFSUri     string `json:"ipfs_uri"`
	Transaction string `json:"transaction"`
}
