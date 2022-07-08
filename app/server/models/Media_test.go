package models

import (
	"testing"
)

func TestFeedRecordArray(t *testing.T) {
	fa := FeedRecordArray{}
	fa = append(fa, FeedRecord{
		Platform: "test",
		ID:       1,
	})
	t.Log(fa.Value())
}
