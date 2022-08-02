package types

type OnChainDispatched struct {
	Platform             string `json:"platform"`
	Username             string `json:"username"`
	FeedID               uint   `json:"feed_id"` // Feed ID in main database
	CrossbellCharacterID string `json:"crossbell_character_id"`
	RawFeed
}

type OnChainRespond struct {
	Platform    string `json:"platform"`
	FeedID      uint   `json:"feed_id"` // Feed ID in main database
	IPFSUri     string `json:"ipfs_uri"`
	Transaction string `json:"tx"`
}
