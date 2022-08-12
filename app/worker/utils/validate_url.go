package utils

import "net/url"

func ValidateUri(uri string) bool {
	if uri == "" {
		return false
	}

	_, err := url.Parse(uri)

	return err == nil

}
