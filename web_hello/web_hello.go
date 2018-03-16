package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<b>Hello World</b><br>")
    //fmt.Fprintf(w, "<b>Hello World</b><br>")
    //fmt.Fprintf(w, "You have requested: %s\n", r.URL.Path)
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}