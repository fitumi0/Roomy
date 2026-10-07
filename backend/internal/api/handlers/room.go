package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

const ROOM_PATH = "/room/{id}"

func HandleRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	regex := `^[a-zA-Z0-9-]+$`
	if !regexp.MustCompile(regex).MatchString(id) {
		http.Error(w, "Invalid room name", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Room ID: %s", id)))
}
