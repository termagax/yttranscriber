package app

import (
	"fmt"
	"net/http"

	"github.com/termagax/yttranscriber/cmd/app/err"
	"github.com/termagax/yttranscriber/cmd/app/format"
	"github.com/termagax/yttranscriber/cmd/app/request"
)

// Transcribe ...
func Transcribe(w http.ResponseWriter, url string) {
	if err.LinkIsValid(url) {
		data := request.GetPage(request.FindInternalURL(url))
		fmt.Fprintf(w, format.JSON(data))
	} else {
		// TODO: handle error
	}
}
