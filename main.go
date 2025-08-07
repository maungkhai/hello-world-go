package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/v1.0/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "pong")
    })

    http.HandleFunc("/v1.0/init", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "init ok")
    })

    addr := "192.168.1.1:8080"
    log.Printf("Starting server at http://%s/v1.0/ping", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal(err)
    }
}
