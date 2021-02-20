package serve

import (
	"log"
	"net/http"
	"strings"

	"github.com/termagax/yttranscriber/cmd/app"
)

func transcriberApp(w http.ResponseWriter, r *http.Request) {
	servePage(w, "templates/transcriber.html", nil, r)
}

func transcript(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()["url"]
	app.Transcribe(w, strings.Join(url, ""))
}

// RunApp ...
func RunApp() {
	http.HandleFunc("/transcriber", transcriberApp)
	http.HandleFunc("/transcriber/transcript", transcript)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
