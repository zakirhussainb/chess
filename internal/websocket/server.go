package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	game "github.com/zakirhussainb/chess/internal/game"
	utils "github.com/zakirhussainb/chess/pkg/utils"
)

var unMappedUsersQueue *utils.Queue
var unMappedGamesQueue *utils.Queue
var chessGame *game.Game
var users []*game.User

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		var msg game.User

		// Parse the JSON into the struct
		err = json.Unmarshal([]byte(message), &msg)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		newUser := game.NewUser(uuid.New().String(), msg.UserName, game.Color(msg.Color), msg.IsPlayer, msg.IsViewer)
		users = append(users, &newUser)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	for {
		// senderMessage := "zak h5"
		// err = conn.WriteMessage(websocket.TextMessage, []byte(senderMessage))
		// if err != nil {
		// 	fmt.Println("Error writing message:", err)
		// 	break
		// }

		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		var msg Message

		// Parse the JSON into the struct
		err = json.Unmarshal([]byte(message), &msg)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		if msg.Type == "init_game" {
			if !unMappedGamesQueue.IsEmpty() {
				chessGame = unMappedGamesQueue.Dequeue().(*game.Game)
				chessGame.Player2 = msg.User
				chessGame.Player2.Color = game.Black
				chessGame.InitBoard()
			} else {
				chessGame = game.NewGame("GA-"+uuid.New().String(), game.Player{}, game.Player{}, []game.Move{})
				chessGame.Player1 = msg.User
				chessGame.Player1.Color = game.White
				unMappedGamesQueue.Enqueue(chessGame)
				// chessGame.Player1 = game.NewPlayer("PL-"+uuid.New().String(), game.White, msg.UserName)
				// unMappedUsersQueue.Enqueue(chessGame.Player1.UserName)
			}

			// unMappedUsersQueue.Enqueue(msg.UserName)
			// if chessGame == nil {
			// 	chessGame = game.NewGame("G_1", game.Player{}, game.Player{}, []game.Move{})
			// 	if chessGame.Player1.UserName == "" {
			// 		chessGame.Player1 = game.NewPlayer("U_1", game.White, msg.UserName)
			// 	} else if chessGame.Player2.UserName == "" {
			// 		chessGame.Player2 = game.NewPlayer("U_2", game.Black, msg.UserName)
			// 	}
			// }

			// // unMappedUsersQueue.Dequeue()
			// if chessGame.Player1.UserName != "" && chessGame.Player2.UserName != "" {
			// 	chessGame.InitBoard()
			// }

		}
	}
}

// func addHandler(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade the HTTP connection to a WebSocket connection
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading:", err)
// 		return
// 	}
// 	defer conn.Close()
// }

func StartServer() {
	unMappedUsersQueue = utils.NewQueue()
	unMappedGamesQueue = utils.NewQueue()

	http.HandleFunc("/play", playHandler)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
