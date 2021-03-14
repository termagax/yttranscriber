package request

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/termagax/yttranscriber/cmd/app/format"
)

// GetPage ...
func GetPage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// TODO: handle error
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: handle error
	}

	return string(html)
}

// FindInternalURL ...
func FindInternalURL(url string) string {
	youtubehtml := GetPage(url)
	regex := regexp.MustCompile(`[^playerCaptionsTracklistRenderer.][^":"]*?(https://www.youtube.com/api/timedtext.*?\")`)
	dataURL := regex.FindString(youtubehtml)
	uri := format.URL(dataURL)
	return uri
}
