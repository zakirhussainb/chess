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

func initBoard(game Game) {
	// Initializing Main Pieces for Player2
	board[0][0] = NewMove(game.Player2, ROOK, "a8", "a8")
	board[0][1] = NewMove(game.Player2, KNIGHT, "b8", "b8")
	board[0][2] = NewMove(game.Player2, BISHOP, "c8", "c8")
	board[0][3] = NewMove(game.Player2, QUEEN, "d8", "d8")
	board[0][4] = NewMove(game.Player2, KING, "e8", "e8")
	board[0][5] = NewMove(game.Player2, BISHOP, "f8", "f8")
	board[0][6] = NewMove(game.Player2, KNIGHT, "g8", "g8")
	board[0][7] = NewMove(game.Player2, ROOK, "h8", "h8")

	// Initializing Pawns for Player2
	board[1][0] = NewMove(game.Player2, PAWN, "a7", "a7")
	board[1][1] = NewMove(game.Player2, PAWN, "b7", "b7")
	board[1][2] = NewMove(game.Player2, PAWN, "c7", "c7")
	board[1][3] = NewMove(game.Player2, PAWN, "d7", "d7")
	board[1][4] = NewMove(game.Player2, PAWN, "e7", "e7")
	board[1][5] = NewMove(game.Player2, PAWN, "f7", "f7")
	board[1][6] = NewMove(game.Player2, PAWN, "g7", "g7")
	board[1][7] = NewMove(game.Player2, PAWN, "h7", "h7")

	// Initializing Main Pieces for Player1
	board[7][0] = NewMove(game.Player1, ROOK, "a1", "a1")
	board[7][1] = NewMove(game.Player1, KNIGHT, "b1", "b1")
	board[7][2] = NewMove(game.Player1, BISHOP, "c1", "c1")
	board[7][3] = NewMove(game.Player1, QUEEN, "d1", "d1")
	board[7][4] = NewMove(game.Player1, KING, "e1", "e1")
	board[7][5] = NewMove(game.Player1, BISHOP, "f1", "f1")
	board[7][6] = NewMove(game.Player1, KNIGHT, "g1", "g1")
	board[7][7] = NewMove(game.Player1, ROOK, "h1", "h1")

	// Initializing Pawns for Player1
	board[6][0] = NewMove(game.Player1, PAWN, "a2", "a2")
	board[6][1] = NewMove(game.Player1, PAWN, "b2", "b2")
	board[6][2] = NewMove(game.Player1, PAWN, "c2", "c2")
	board[6][3] = NewMove(game.Player1, PAWN, "d2", "d2")
	board[6][4] = NewMove(game.Player1, PAWN, "e2", "e2")
	board[6][5] = NewMove(game.Player1, PAWN, "f2", "f2")
	board[6][6] = NewMove(game.Player1, PAWN, "g2", "g2")
	board[6][7] = NewMove(game.Player1, PAWN, "h2", "h2")
}

func Init() {
	fmt.Println("Game started")
	player1 = NewPlayer("P_1", White, "zakir")
	player2 = NewPlayer("P_2", Black, "thahsin")
	var moves []Move
	game = NewGame("G_1", player1, player2, moves)
	initBoard(game)

	// printBoard()

}

// func printBoard() {
// 	for i := range 8 {
// 		for j := range 8 {
// 			move := board[i][j]
// 			if move.PlayerID != "" {
// 				fmt.Printf("%s:%s ", move.PlayerID, move.From)
// 			} else {
// 				fmt.Print(". ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
