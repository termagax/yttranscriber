package main

import (
	"fmt"
	"yttranscript/jsonops"
	"yttranscript/request"
)

func main() {
	regex := ""
	url := "https://www.youtube.com/watch?v=62dl2TxGQYw"
	uri := request.ParseHTML(url, regex)

	// find the URI with the use of regex plz
	fmt.Println(uri)
	jsonops.GetFromURI(uri)
	jsonops.Parse()

}
