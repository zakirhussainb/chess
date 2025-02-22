package game

type Move struct {
	PlayerID string
	Piece    PIECE
	From     string
	To       string
}

func NewMove(playerID string, piece PIECE, from, to string) Move {
	return Move{
		PlayerID: playerID,
		Piece:    piece,
		From:     from,
		To:       to,
	}
}
