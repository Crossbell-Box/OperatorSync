package types

type OnChainRequest struct {
	Platform             string `json:"platform"`
	Username             string `json:"username"`
	FeedID               uint   `json:"feed_id"` // Feed ID in main database
	CrossbellCharacterID string `json:"crossbell_character_id"`
	RawFeed
}

type OnChainResponse struct {
	IsSucceeded bool   `json:"is_succeeded"`
	Message     string `json:"message"`
	Platform    string `json:"platform"`
	FeedID      uint   `json:"feed_id"` // Feed ID in main database
	IPFSUri     string `json:"ipfs_uri"`
	Transaction string `json:"tx"`
	CharacterID int64  `json:"character_id"`
	NoteID      int64  `json:"note_id"`
}
