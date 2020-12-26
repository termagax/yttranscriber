package request

import "regexp"

func format(uri string) string {

	newRegex := regexp.MustCompile("[^playerCaptionsTracklistRenderer.*?](https://www.youtube.com/api/timedtext.*?\")")

	uri = string(newRegex.FindString(uri))

	return uri
}
