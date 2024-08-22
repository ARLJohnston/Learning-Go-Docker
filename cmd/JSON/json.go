package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
		"dateValue": time.Date(2022, 3, 2, 9, 10, 0, 0, time.Local),
		"nullVal":   nil,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json")
		return
	}

	fmt.Printf("%s\n", jsonData)
}
