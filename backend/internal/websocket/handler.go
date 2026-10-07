package websocket

import (
	"net/http"

	"github.com/coder/websocket"
)

// TODO:

// here upgrade handler
const WS_PATH = "/ws"

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	websocket.Accept(w, r, nil)
}
