package request

import (
	"regexp"
)

// ParseHTML finds URI
func ParseHTML(url string, regex string) string {
	vars := uriFinderComponent{
		uri:   "",
		data:  get(url),
		regex: regexp.MustCompile(regex),
	}

	vars.uri = string(vars.regex.FindString(vars.data))
	return vars.uri
}
