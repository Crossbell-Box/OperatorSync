package models

import (
	"testing"
)

func TestFeedRecordArray(t *testing.T) {
	fa := MediaFeedRecordArray{}
	fa = append(fa, MediaFeedRecord{
		Platform: "test",
		ID:       1,
	})
	t.Log(fa.Value())
}
