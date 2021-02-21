package request

import (
	"io/ioutil"
	"net/http"
)

// GetPage ...
func GetPage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// TODO: do something with this error
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// do something else lol
	}

	return string(html)
}
