package main

import (
    "log"
    "net/http"
    "github.com/andrewwhiteside/continuum-core/pkg/api"
)

func main() {
    mux := http.NewServeMux()
    api.RegisterHandlers(mux)

    log.Println("Starting Continuum v0.1 on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}