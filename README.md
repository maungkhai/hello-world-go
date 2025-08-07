# Go WebSocket Chat Application

A simple real-time chat application built with Go and WebSocket. This project includes both a chat server and a web-based UI.

## Features

- Real-time messaging using WebSocket
- Simple and clean web interface
- Multiple concurrent chat clients support
- Basic API endpoints (/v1.0/ping and /v1.0/init)

## Requirements

- Go 1.21 or later
- Gorilla WebSocket package

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```
2. Open your web browser and navigate to:
   ```
   http://localhost:8080/v1.0/chat
   ```

## API Endpoints

- `/v1.0/chat` - Web chat interface
- `/v1.0/ws` - WebSocket endpoint for chat
- `/v1.0/ping` - Health check endpoint
- `/v1.0/init` - Initialization endpoint

## Development

The project uses:
- `github.com/gorilla/websocket` for WebSocket functionality
- Standard Go HTTP server for serving static files and API endpoints

Feel free to contribute or modify for your needs!
