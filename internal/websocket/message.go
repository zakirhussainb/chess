package websocket

import (
	game "github.com/zakirhussainb/chess/internal/game"
)

type Message struct {
	Type string    `json:"type"`
	User game.User `json:"user"`
}
