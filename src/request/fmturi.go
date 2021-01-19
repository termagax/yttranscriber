package request

import (
	"regexp"
	"strings"
)

func format(uri string) string {

	//
	var slice []string
	newRegex := regexp.MustCompile("[^playerCaptionsTracklistRenderer.*?](https://www.youtube.com/api/timedtext.*?\")")

	uri = string(newRegex.FindString(uri))

	slice = strings.Split(uri, "\\u0026")
	uri = strings.Join(slice, "/")

	slice = strings.Split(uri, "\"")
	uri = strings.Join(slice, "")

	return uri
}
