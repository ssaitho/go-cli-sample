package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/auth", authHandler)
    http.HandleFunc("/hello", helloHandler)
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
