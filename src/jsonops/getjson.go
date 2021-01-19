package jsonops

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Number int `json:"word"`
}

// GetFromURI gets the json from the given URI
func GetFromURI(uri string) {

	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`

	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println("could not ")
		return
	}
	fmt.Println(people1.Number)
}
