package websocket

import "github.com/coder/websocket"

type Client struct {
	conn   *websocket.Conn // The socket connection
	send   chan []byte     // Buffered channel for outbound messages
	roomID string
}
