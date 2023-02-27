package types

type NoteAttachment struct {
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	Content  string `json:"content,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize uint   `json:"size_in_bytes"`

	// Image & Video specified
	Alt    string `json:"alt,omitempty"`
	Width  uint   `json:"width,omitempty"`
	Height uint   `json:"height,omitempty"`
}

type NoteAttributes struct {
	Value       string `json:"value"` // string | number | boolean | null
	TraitType   string `json:"trait_type,omitempty"`
	DisplayType string `json:"display_type"` // 'string' | 'number' | 'date' | 'boolean'
}

type NoteMetadata struct {
	Type           string           `json:"type"` // "note"
	Authors        []string         `json:"authors"`
	Title          string           `json:"title"`
	Content        string           `json:"content"`
	Tags           []string         `json:"tags"`
	Attachments    []NoteAttachment `json:"attachments"`
	Sources        []string         `json:"sources"`
	Attributes     []NoteAttributes `json:"attributes,omitempty"`
	ExternalUrls   []string         `json:"external_urls,omitempty"`
	ContentWarning string           `json:"content_warning,omitempty"` // 'nsfw' | 'sensitive' | 'spoiler'
	DatePublished  string           `json:"date_published,omitempty"`  // ISO8601
}
