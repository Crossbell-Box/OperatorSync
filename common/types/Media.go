package types

type Media struct {
	OriginalURI string `json:"original_uri"`
	IPFSUri     string `json:"ipfs_uri" gorm:"index;column:ipfs_uri"`
	FileSize    uint   `json:"file_size"`
	ContentType string `json:"content_type"`
}
