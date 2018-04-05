//http://polyglot.ninja/golang-json/

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	jb := Person{
		Name:   "Jim Beam",
		Age:    103,
		Emails: []string{"test@gmail.com", "test2@yahoo.com"},
	}
	json_bytes, _ := json.Marshal(jb)
	fmt.Println(string(json_bytes))
}
