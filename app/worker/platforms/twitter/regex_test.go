package twitter

import (
	"strings"
	"testing"
)

func TestImageRegex(t *testing.T) {
	rawContent := "图片发送测试 2210111430<br><img style=\"\" src=\"https://pbs.twimg.com/media/FexLzY-UUAAUFs_?format=jpg&amp;name=orig\" referrerpolicy=\"no-referrer\">"

	var medias []string
	imgs := imageRegex.FindAllStringSubmatch(rawContent, -1)
	for _, img := range imgs {
		medias = append(medias, img[1])
		rawContent = strings.ReplaceAll(rawContent, img[0], "")
	}

	t.Log(medias)
	t.Log(rawContent)
}

func TestVideoRegex(t *testing.T) {
	rawContent := "Elegant!<br><br>Crossbell: You can now enjoy Crossbell on @raycastapp! You can:<br>-browse your personal feed<br>-view the latest notes<br>-search across the network<br>-open in your browser<br>All happens in a flash!<br><br><video src=\"https://video.twimg.com/ext_tw_video/1575125440624267268/pu/vid/1280x720/8oafB-WK57me-cos.mp4?tag=12\" controls=\"controls\" poster=\"https://pbs.twimg.com/ext_tw_video_thumb/1575125440624267268/pu/img/IPEscFVyC1zXsWvq.jpg\"></video> <a href=\"https://www.raycast.com/Songkeys/crossbell\" target=\"_blank\" rel=\"noopener noreferrer\">https://www.raycast.com/Songkeys/crossbell</a>"

	var medias []string
	videos := videoRegex.FindAllStringSubmatch(rawContent, -1)
	for _, video := range videos {
		medias = append(medias, video[1])
		rawContent = strings.ReplaceAll(rawContent, video[0], "")
	}

	t.Log(medias)
	t.Log(rawContent)
}
