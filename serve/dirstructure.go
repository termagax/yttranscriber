package serve

import (
	"html/template"
	"log"
	"net/http"
	"transcriber/app"
)

func servePage(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func transcriberApp(w http.ResponseWriter, r *http.Request) {
	servePage(w, "transcriber.html", nil)
}

func transcript(w http.ResponseWriter, r *http.Request) {
	app.Transcribe(w, "https://www.youtube.com/watch?v=vsTTXYxydOE")
}

// RunApp ...
func RunApp() {
	http.HandleFunc("/transcriber", transcriberApp)
	http.HandleFunc("/transcriber/transcript", transcript)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
