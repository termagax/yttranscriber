package app

import (
	"fmt"
	"net/http"
)

// Transcribe ...
func Transcribe(w http.ResponseWriter, url string) {
	fmt.Fprintln(w, "no matter what you input here, it will not make a difference")
}
