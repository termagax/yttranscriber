package app

import (
	"fmt"
	"net/http"

	"github.com/termagax/yttranscriber/cmd/app/request"
)

// Transcribe ...
func Transcribe(w http.ResponseWriter, url string) {
	if request.LinkIsValid(url) {

		raw := request.GetPage(url)
		fmt.Fprintf(w, raw)
	}
}
