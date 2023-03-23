package utils

import "regexp"

var (
	TrimRegex *regexp.Regexp
)

func init() {
	TrimRegex = regexp.MustCompile("[ \t\r\n\f]") // Should just be \s, while it seems that Go doesn't support this
}

func StringPointerOmitEmpty(str string) *string {
	trimmedStr := TrimRegex.ReplaceAllString(str, "")
	if trimmedStr == "" {
		return nil
	} else {
		return &str // Original String
	}
}
