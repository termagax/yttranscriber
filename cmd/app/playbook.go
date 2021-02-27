package app

import (
	"regexp"

	"github.com/termagax/yttranscriber/cmd/app/format"
	"github.com/termagax/yttranscriber/cmd/app/request"
)

func finalTranscript(url string) string {
	var response string
	youtubehtml := request.GetPage(url)
	// TODO: figure out this regex
	regex := regexp.MustCompile("")
	dataUrl := request.FindPatternWithinString(youtubehtml, regex)

	// get data url to be in string format and unmarshaled (if any)
	uri := format.URL(dataUrl)

	dataResponse := request.GetPage(uri)

	// unmarshal and format json or xml request
	if true {
		response = format.JSON(dataResponse, []string{"word"})
	} else if false {
		response = format.XML(dataResponse, []string{"hello"})
	} else {
		// TODO: handle error
	}
	return response
}
