package app

import (
	"fmt"
	"net/http"
)

// Transcribe ...
func Transcribe(w http.ResponseWriter, url string) {
	fmt.Fprintln(w, url)
	fmt.Println(url)
}
