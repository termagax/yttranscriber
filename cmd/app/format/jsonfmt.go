package format

import (
	"encoding/json"
	"fmt"
)

// JSON ...
func JSON(raw string) string {
	var transcript string

	var junkNav map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &junkNav); err != nil {
		panic(err)
	}

	getEventsOutOfJunk := junkNav["events"].([]interface{})

	for _, v := range getEventsOutOfJunk {
		getSegsOutOfEvents := v.(map[string]interface{})
		if getSegsOutOfEvents["segs"] != nil {
			lol := getSegsOutOfEvents["segs"]
			list := lol.([]interface{})
			for _, i := range list {
				netlock := i.(map[string]interface{})
				s := fmt.Sprintf("%v", netlock["utf8"])
				transcript += s
			}
		}
	}

	return transcript

}
