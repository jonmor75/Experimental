package main

import (
	"fmt"
	"net/http"
)

//An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
//An http.Request is a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

/*
func main() {
	http.HandleFunc("/", handler)
	// ListenAndServe always returns an error, since it only returns when an unexpected error occurs. In order to log that error we wrap the function call with log.Fatal.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
