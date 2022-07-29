package types

type NoteMetadata struct {
	Type        string   `json:"type"` // "note"
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	Attachments []struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Content  string `json:"content"`
		MimeType string `json:"mime_type"`
		FileSize string `json:"size_in_bytes"`
	} `json:"attachments"`
	Sources []string `json:"sources"`
}

type NoteOnChain struct {
	ContentURI string `json:"contentURI,omitempty"`
}
