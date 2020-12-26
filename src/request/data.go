package request

import "regexp"

type uriFinderComponent struct {
	uri, data string
	regex     *regexp.Regexp
}
