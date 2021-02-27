package app

import (
	"fmt"
	"net/http"

	"github.com/termagax/yttranscriber/cmd/err"
)

// Transcribe ...
func Transcribe(w http.ResponseWriter, url string) {
	if err.LinkIsValid(url) {
		fmt.Fprintf(w, finalTranscript(url))
	} else {
		// TODO: handle error
	}
}
