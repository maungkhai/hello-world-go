package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/gorilla/websocket"
)

func BenchmarkWebSocketChat(b *testing.B) {
    // Start the chat server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        upgrader := websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool { return true },
        }
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            b.Fatalf("Failed to upgrade connection: %v", err)
        }
        defer conn.Close()

        for {
            _, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }
            // Echo back the message
            err = conn.WriteMessage(websocket.TextMessage, msg)
            if err != nil {
                return
            }
        }
    }))
    defer server.Close()

    url := "ws" + strings.TrimPrefix(server.URL, "http")

    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        // Create WebSocket client
        ws, _, err := websocket.DefaultDialer.Dial(url, nil)
        if err != nil {
            b.Fatalf("Failed to connect: %v", err)
        }
        defer ws.Close()

        msg := []byte("benchmark message")
        for pb.Next() {
            // Send message
            err := ws.WriteMessage(websocket.TextMessage, msg)
            if err != nil {
                b.Fatalf("Write failed: %v", err)
            }

            // Read response
            _, response, err := ws.ReadMessage()
            if err != nil {
                b.Fatalf("Read failed: %v", err)
            }

            if string(response) != string(msg) {
                b.Fatalf("Expected %s, got %s", msg, response)
            }
        }
    })
}