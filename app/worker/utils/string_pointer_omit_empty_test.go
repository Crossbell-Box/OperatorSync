package utils

import "testing"

func TestStringPointerOmitEmpty(t *testing.T) {
	t.Log(StringPointerOmitEmpty("") == nil)                  // Just empty
	t.Log(StringPointerOmitEmpty("  ") == nil)                // Spaces
	t.Log(StringPointerOmitEmpty("\n\n\r\n") == nil)          // New lines
	t.Log(StringPointerOmitEmpty("\t\t\n  ") == nil)          // Mixed
	t.Log(StringPointerOmitEmpty("\n\n123\rabc\nfoo") != nil) // With value
}
