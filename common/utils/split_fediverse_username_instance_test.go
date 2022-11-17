package utils

import "testing"

func TestSplitFediverseUsernameInstance(t *testing.T) {
	usernameAtInstance := "candinya@thenomads.social"
	t.Log(SplitFediverseUsernameInstance(usernameAtInstance))
}
