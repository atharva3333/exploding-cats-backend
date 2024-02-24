package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
    })

    fmt.Println("Listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
