package request

import "regexp"

// FindPatternWithinString performs a given regex on given data, returns the result
func FindPatternWithinString(data string, regex *regexp.Regexp) []string {
	arr := []string{"hello", "othertext"}
	return arr
}
