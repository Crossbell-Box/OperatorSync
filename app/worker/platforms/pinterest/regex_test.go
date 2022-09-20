package pinterest

import (
	"strings"
	"testing"
)

func TestPinterestFeedRegex(t *testing.T) {

	rawContent := "<a href=\"https://www.pinterest.com/pin/870742909180165305/\">\n                  <img src=\"https://i.pinimg.com/236x/03/07/6d/03076d3473fa18e15cd051fa22dc2dbf.jpg\"></a>\n                  Nya Avatar 😘"

	var originalSizeImgs []string
	anchors := anchorRegex.FindAllString(rawContent, -1)
	for _, a := range anchors {
		imgs := imageRegex.FindAllStringSubmatch(a, -1)
		for _, img := range imgs {
			originalsUri := strings.Replace(img[1], "/236x/", "/originals/", 1)
			originalSizeImgs = append(originalSizeImgs, originalsUri)
			rawContent = strings.Replace(rawContent, a, strings.Replace(img[0], img[1], originalsUri, 1), 1)
		}
	}

	t.Log(rawContent)
	t.Log(originalSizeImgs)
}
