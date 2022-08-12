package utils

import (
	"testing"
)

func TestValidateUri(t *testing.T) {

	if ValidateUri("") {
		t.Fail()
	}
	if !ValidateUri("https://candinya.com/") {
		t.Fail()
	}
}
