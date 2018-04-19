package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mapToConvert := map[string]int{"jobid": 1, "error": 5}
	mapToConvertJSON, err := json.Marshal(mapToConvert)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mapToConvertJSON))
}
