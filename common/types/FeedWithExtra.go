package types

import (
	"github.com/mmcdole/gofeed"
	"time"
)

type ExtraLinks struct {
	URL         string `json:"url"`
	Type        string `json:"type"` // reply / repost / quote
	ContentHTML string `json:"content_html"`
}

type ExtraSpec struct {
	Links []ExtraLinks `json:"links"`
}

// Modified from gofeed

// FeedWithExtra is the universal Feed type that atom.Feed
// and rss.Feed gets translated to. It represents
// a web feed.
// Sorting with sort.Sort will order the Items by
// oldest to newest publish time.
type FeedWithExtra struct {
	gofeed.Feed

	Items []*ItemWithExtra `json:"items"`
}

// ItemWithExtra is the universal Item type that atom.Entry
// and rss.Item gets translated to.  It represents
// a single entry in a given feed.
type ItemWithExtra struct {
	gofeed.Item

	ID            string    `json:"id"`
	URL           string    `json:"url"`
	DatePublished time.Time `json:"date_published"`
	ContentHTML   string    `json:"content_html"`
	Extra         ExtraSpec `json:"_extra"`
}
