package serve

import (
	"fmt"
	"html/template"
	"net/http"
)

func servePage(w http.ResponseWriter, filename string, data interface{}, r *http.Request) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Printf("FAIL: cannot find %s!", filename)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
