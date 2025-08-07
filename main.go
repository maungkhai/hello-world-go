package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
    "github.com/gorilla/websocket"
)

func main() {
    http.HandleFunc("/v1.0/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "pong")
    })

    http.HandleFunc("/v1.0/init", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "init ok")
    })

    // Serve chat UI
    http.HandleFunc("/v1.0/chat", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    // WebSocket chat handler
    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool { return true },
    }
    var clients = make(map[*websocket.Conn]bool)
    var mu sync.Mutex

    http.HandleFunc("/v1.0/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            return
        }
        mu.Lock()
        clients[conn] = true
        mu.Unlock()
        defer func() {
            mu.Lock()
            delete(clients, conn)
            mu.Unlock()
            conn.Close()
        }()
        for {
            _, msg, err := conn.ReadMessage()
            if err != nil {
                break
            }
            mu.Lock()
            for c := range clients {
                c.WriteMessage(websocket.TextMessage, msg)
            }
            mu.Unlock()
        }
    })

    addr := "localhost:8080"
    log.Printf("Starting server at http://%s/v1.0/chat", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal(err)
    }
}
