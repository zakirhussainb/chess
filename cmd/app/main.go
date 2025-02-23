package main

import (
	"fmt"

	"github.com/zakirhussainb/chess/internal/websocket"
)

func main() {
	fmt.Println("Starting the game...")
	websocket.StartServer()
	// game.Init()
}
