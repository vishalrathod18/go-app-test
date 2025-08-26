package main

import (
    "fmt"
    "net/http"
    "go-app/internal/handler"
)

func main() {
    http.HandleFunc("/", handler.Home)
    fmt.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}