package websocket

type Message struct {
	RoomID  string `json:"room_id"`
	Content []byte `json:"content"` // TODO: structurize
}
