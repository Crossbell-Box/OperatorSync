package types

import (
	"github.com/lib/pq"
	"time"
)

// RawFeed : Minimal content of feed (just for work response)
type RawFeed struct {
	Language       string         `json:"language"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Content        string         `json:"text"`
	Link           string         `json:"link"`
	UpdatedAt      time.Time      `json:"updated_at"`
	PublishedAt    time.Time      `json:"published_at"`
	Authors        pq.StringArray `json:"authors" gorm:"type:text[]"`
	GUID           string         `json:"guid"`
	Image          string         `json:"image"`
	Categories     pq.StringArray `json:"categories" gorm:"type:text[]"`
	Media          []Media        `json:"media" gorm:"-"`
	ContentWarning string         `json:"content_warning"` // 'nsfw' | 'sensitive' | 'spoiler'
}
