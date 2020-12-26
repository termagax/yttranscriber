package main

import (
	"fmt"
	"yttranscript/jsonops"
	"yttranscript/request"
)

func main() {

	regex := "playerCaptionsTracklistRenderer.*?(https://www.youtube.com/api/timedtext.*?\")"
	//regex := "/playerCaptionsTracklistRenderer.*?(youtube.com/api/timedtext.*?)\"/"
	url := "https://www.youtube.com/watch?v=r7SO-Oq3d5E&t=1128s"
	uri := request.ParseHTML(url, regex)

	// find the URI with the use of regex plz
	fmt.Println(uri)
	jsonops.GetFromURI(uri)
	jsonops.Parse()

}
