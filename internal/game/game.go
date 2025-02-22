package game

import "fmt"

var board [8][8]Move
var player1 Player
var player2 Player
var game Game

type Game struct {
	ID      string
	Player1 Player
	Player2 Player
	Moves   []Move
}

func NewGame(id string, player1, player2 Player, moves []Move) Game {
	return Game{
		ID:      id,
		Player1: player1,
		Player2: player2,
		Moves:   moves,
	}
}

func init_board(game Game) {
	board[0][0] = NewMove(game.Player2.ID, ROOK, "a8", "a8")
}

// func play_game(gameID, p1_ID, p2_ID int, p1_piece, p2_piece PIECE) {

// }

func Start() {
	fmt.Printf("Game started")
	player1 = NewPlayer("P_1", "white", "zakir")
	player2 = NewPlayer("P_2", "black", "thahsin")
	var moves []Move
	game = NewGame("G_1", player1, player2, moves)
	init_board(game)
}
