package types

import (
	"github.com/lib/pq"
	"time"
)

//type Media struct {
//	Url    string `json:"url"`
//	Length uint   `json:"length"`
//	Type   string `json:"type"`
//}
//
//type MediaArray []Media
//
//func (ma *MediaArray) Scan(src interface{}) error {
//	return json.Unmarshal(src.([]byte), ma)
//}
//
//func (ma *MediaArray) Value() (driver.Value, error) {
//	val, err := json.Marshal(ma)
//	return string(val), err
//}

// RawFeed : Minimal content of feed (just for work response)
type RawFeed struct {
	Language    string         `json:"language"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Content     string         `json:"text"`
	Link        string         `json:"link"`
	UpdatedAt   time.Time      `json:"updated_at"`
	PublishedAt time.Time      `json:"published_at"`
	Authors     pq.StringArray `json:"authors"`
	GUID        string         `json:"guid"`
	Image       string         `json:"image"`
	Categories  pq.StringArray `json:"categories"`
	Media       pq.StringArray `json:"media"`
	//Media       MediaArray     `json:"media" gorm:"type:longtext"` // Enclosure
}
