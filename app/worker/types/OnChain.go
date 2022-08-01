package types

type NoteAttachment struct {
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	Content  string `json:"content,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize uint   `json:"size_in_bytes"`
}

type NoteMetadata struct {
	Type        string           `json:"type"` // "note"
	Title       string           `json:"title"`
	Content     string           `json:"content"`
	Tags        []string         `json:"tags"`
	Attachments []NoteAttachment `json:"attachments"`
	Sources     []string         `json:"sources"`
}

type NoteOnChain struct {
	ContentURI string `json:"contentURI,omitempty"`
}
