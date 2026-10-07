package api

import (
	"roomy/internal/api/handlers"
	"roomy/internal/websocket"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// TODO: internal.HandleWebSocket
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get(handlers.ROOM_PATH, handlers.HandleRoom)

	r.Get(websocket.WS_PATH, websocket.HandleWebSocket)

	return r
}
