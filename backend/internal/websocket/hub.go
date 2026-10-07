package websocket

type Hub struct {
	rooms      map[string]map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func (h *Hub) Run() {
	for {
	}
}
