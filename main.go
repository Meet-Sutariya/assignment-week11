package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/about", about)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the API Server!")
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a simple API built with Go.")
}
