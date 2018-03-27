//http://polyglot.ninja/golang-making-http-requests/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	GetRequest()
	//PostRequest()
	//UrlPost()
}

// GetRequest is used to create a HTTP Get Request
func GetRequest() {
	resp, err := http.Get("https://httpbin.org/get")
	//resp, err := http.Get("http://devicecheck.scapp-corp.swisscom.com")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	defer resp.Body.Close()
}

// PostRequest is used to create a HTTP Post Request
func PostRequest() {

	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
	defer resp.Body.Close()
}

// UrlPost is used to create a URL Post (Form) Request
func UrlPost() {

	formData := url.Values{
		"name": {"jonathan"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	//log.Println(result)
	log.Println(result["form"])
	defer resp.Body.Close()

}
