package game

type PIECE string

const (
	PAWN   PIECE = "Pawn"
	KNIGHT PIECE = "Knight"
	BISHOP PIECE = "Bishop"
	ROOK   PIECE = "Rook"
	QUEEN  PIECE = "Queen"
	KING   PIECE = "King"
)

type Color int

const (
	White Color = iota
	Black
)
