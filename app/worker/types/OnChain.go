package types

type NoteAttachment struct {
	Name     *string `json:"name,omitempty"`
	Address  *string `json:"address,omitempty"`
	Content  *string `json:"content,omitempty"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *uint   `json:"size_in_bytes,omitempty"`

	// Image & Video specified
	Alt    *string `json:"alt,omitempty"`
	Width  *uint   `json:"width,omitempty"`
	Height *uint   `json:"height,omitempty"`
}

type NoteAttributes struct {
	Value       string `json:"value"` // string | number | boolean | null
	TraitType   string `json:"trait_type,omitempty"`
	DisplayType string `json:"display_type"` // 'string' | 'number' | 'date' | 'boolean'
}

// https://github.com/Crossbell-Box/crossbell.js/blob/c5ff4c2/src/types/metadata/note.ts#L49
type NoteMetadata struct {
	Type           string           `json:"type"` // "note"
	Authors        []string         `json:"authors"`
	Title          *string          `json:"title,omitempty"`
	Content        *string          `json:"content,omitempty"`
	Tags           []string         `json:"tags,omitempty"`
	Attachments    []NoteAttachment `json:"attachments,omitempty"`
	Sources        []string         `json:"sources,omitempty"`
	Attributes     []NoteAttributes `json:"attributes,omitempty"`
	ExternalUrls   []string         `json:"external_urls,omitempty"`
	ContentWarning *string          `json:"content_warning,omitempty"` // 'nsfw' | 'sensitive' | 'spoiler'
	DatePublished  *string          `json:"date_published,omitempty"`  // ISO8601
}
