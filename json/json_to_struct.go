//http://polyglot.ninja/golang-json/

package main

import (
	"encoding/json"
	"fmt"
)

//Person ....
type Person struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Emails []string `json:"emails"`
}

func main() {
	json_bytes := []byte(`
		{
			"Name" : "Jim Beam",
			"Age" : 42,
			"Emails" : ["test@gmail.com","test@yahoo.com"]
		}
	`)
	jb := Person{}
	err := json.Unmarshal(json_bytes, &jb)
	if err != nil {
		panic(err)
	}
	fmt.Println(jb.Name, jb.Age, jb.Emails)
}
