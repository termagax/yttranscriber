package format

import (
	"strings"
)

// URL ...
func URL(url string) string {
	appended := "&kind=asr&lang=en&fmt=json3&xorb=2&xobt=3&xovt=3"
	unmarshal := strings.Split(url, "\\u0026")
	unmarshaled := strings.Join(unmarshal, "&")
	removeQuote := strings.Split(unmarshaled, "\"")
	quoteRemoved := strings.Join(removeQuote, "")
	return quoteRemoved + appended
}
