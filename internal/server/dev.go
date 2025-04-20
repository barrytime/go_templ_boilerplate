package server

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	// WebSocket upgrader
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Store connected clients
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

// NotifyClients sends a reload signal to all connected WebSocket clients
func NotifyClients() {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte("reload"))
		if err != nil {
			log.Printf("WebSocket error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func HandleHotReload(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return err
	}
	defer conn.Close()

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	// Keep connection alive until it closes
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()
			return nil
		}
	}
}
