package app

import (
	"fmt"
	"regexp"

	"github.com/termagax/yttranscriber/cmd/app/format"
	"github.com/termagax/yttranscriber/cmd/app/request"
)

func finalTranscript(url string) string {
	var response string
	youtubehtml := request.GetPage(url)
	regex := regexp.MustCompile(`[^playerCaptionsTracklistRenderer.][^":"]*?(https://www.youtube.com/api/timedtext.*?\")`)
	dataURL := regex.FindString(youtubehtml)
	fmt.Println(dataURL)

	// get data url to be in string format and unmarshaled (if any)
	uri := format.URL(dataURL)
	fmt.Println(uri)

	//dataResponse := request.GetPage(uri)

	// // unmarshal and format json or xml request
	// if true {
	// 	response = format.JSON(dataResponse, []string{"word"})
	// } else if false {
	// 	response = format.XML(dataResponse, []string{"hello"})
	// } else {
	// 	// TODO: handle error
	// }
	return response
}
