package mastodon

import "testing"

func TestCustomEmojisRegex(t *testing.T) {
	rawText := `<p>Contabo 美国又炸了，实在离谱 <img rel="emoji" draggable="false" width="16" height="16" class="emojione" style="width: 1.1em; height: 1.1em; object-fit: contain; vertical-align: middle; margin: -.2ex .15em .2ex" alt=":blobcatgooglytrash:" title=":blobcatgooglytrash:" src="https://files.eihei.net/custom_emojis/images/000/021/683/original/a9b36e44a6b29b5e.png"></p>`

	results := customEmojisRegex.FindAllStringSubmatch(rawText, -1)

	for _, r := range results {
		t.Log(r[1])
	}
}
