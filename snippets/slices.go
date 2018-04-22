package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Creating a cpe list
	type cpe struct {
		id  int
		url string
	}
	var cpeList []cpe
	for i := 0; i < 5; i++ {
		cpeList = append(cpeList, cpe{i, ""})
	}
	// fmt.Println(cpeList)

	//for each cpe in the list add an url
	for i, v := range cpeList {
		url := "https://fakeurl/"
		url = url + strconv.Itoa(v.id)
		// v.url = url // This does not work
		cpeList[i].url = url
	}
	fmt.Println(cpeList)
}
