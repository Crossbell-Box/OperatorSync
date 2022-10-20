package y2b_channel

import "testing"

func TestY2bChannelDescriptionRegex(t *testing.T) {
	demoY2bChannelPage := "" // Disabled cause it's too long

	demoY2bChannelDescription := y2bChannelDescriptionRegex.FindStringSubmatch(demoY2bChannelPage)

	t.Log(demoY2bChannelDescription[1])
}
