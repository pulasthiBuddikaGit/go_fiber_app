package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}